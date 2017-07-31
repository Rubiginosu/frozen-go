package dmserver

import (
	"fmt"
	"os/exec"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

// 按照错误码准备环境
func (server *ServerLocal) EnvPrepare() error {
	serverDataDir := "../servers/server" + strconv.Itoa(server.ID) // 在一开头就把serverDir算好，增加代码重用
	// 文件夹不存在则创建文件夹
	fileInfo, err := os.Stat(serverDataDir + "/serverData")
	if err != nil {
		err := server.prepareDir(serverDataDir)
		if err != nil {
			return err
		}
	} else if fileInfo.Mode() != 0664 {
		os.Chown(serverDataDir, 0, config.DaemonServer.UserId)
		os.Chmod(serverDataDir, 0664)
	} // 创建后判断并设置目录权限
	// 得到执行conf
	execConf, err0 := server.loadExecutableConfig()
	if err0 != nil {
		return err0
	}

	if _, err0 := os.Stat(serverDataDir + ".loop"); err0 != nil {
		fmt.Println("No loop file found!")
		//  新增 loop
		cmd := exec.Command("/bin/dd", "if=/dev/zero", "bs=1024",
			"count="+strconv.Itoa(execConf.MaxHardDisk), "of=../servers/server"+strconv.Itoa(server.ID)+".loop")
		fmt.Print("Writing file...")
		err := cmd.Run()
		if err != nil {
			return err
		}
		fmt.Println("Done")

		fmt.Println("Formatting...")
		cmd2 := exec.Command("/sbin/mkfs.ext4", serverDataDir+".loop")
		err2 := cmd2.Run()
		fmt.Println("Done")
		if err2 != nil {
			return err2
		}

	}
	os.MkdirAll(serverDataDir+"/lib", 0664)
	if _,err := os.Stat("/lib64");err == nil{ // 32位系统貌似没有lib64,那就不新建了

		os.MkdirAll(serverDataDir+"/lib64", 0664)	// 之所以不用GOARCH 是因为可能这个软件是32位，但运行着的系统或许是64位呢？
		// 这个谁说的准？ 哈哈～
	}

	cmd := exec.Command("/bin/mount", "-o", "loop", serverDataDir+".loop", serverDataDir)
	err2 := cmd.Run()
	if err != nil {
		return err2
	}

	/////////////////////////////////////////////////////////////////
	return nil
}

func (server *ServerLocal) prepareDir(serverDataDir string) error {
	err := os.MkdirAll(serverDataDir, 664)
	return err
}

func (server *ServerLocal) loadExecutableConfig() (ExecConf, error) {
	var newServerRuntimeConf ExecConf
	b, err := ioutil.ReadFile("../exec/" + server.Executable + ".json") // 将配置文件读入
	if err != nil {
		// 若在读文件时就有异常则停止反序列化
		return newServerRuntimeConf, err
	}
	err2 := json.Unmarshal(b, &newServerRuntimeConf) //使用自带的json库对读入的东西反序列化
	if err2 != nil {
		return newServerRuntimeConf, err
	}
	return newServerRuntimeConf, nil // 返回结果
}

func (server *ServerLocal) mountDirs() error{
	serverDataDir := "../servers/server" + strconv.Itoa(server.ID) // 在一开头就把serverDir算好，增加代码重用
	execConfig, err := server.loadExecutableConfig()
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/mount -o bind /lib " + serverDataDir + "/lib")
	cmd.Run()
	if _,err := os.Stat("/lib64");err == nil{
		// 这里不用serverDataDir是处于安全考虑，万一小天才给我在../新建了一个lib64 那我把没有的lib64挂载过来就纯属多此一举了
		cmd  := exec.Command("/bin/mount -o bind /lib64 " + serverDataDir + "/lib64")
		cmd.Run()
	}
	cmd2 := exec.Command(generateMountCommand(execConfig.Mount,serverDataDir))
	cmd2.Run()
	return nil
}

/*
如果本来目录都没有当然不予挂载，有可能32位操作系统让挂/usr/lib64，这里需要鲁棒一下
如果原目录有，目的目录没有，则新建再挂载;如果原目录也没有，则不管，跳过
如果原目录和新目录都有，则直接挂载
Example : dirs : {"/bin","/usr/bin"}
Return : "/bin/mount -o bind /bin $serverDataDir/bin;/bin/mount /usr/bin $serverDataDir/usr/bin;"
 */
func generateMountCommand(dirs []string,serverDataDir string) string{
	var toBeMounted []string
	var commands string
	for i:=0;i<len(dirs);i++{
		if _,err := os.Stat(dirs[i]);err == nil {
			toBeMounted = append(toBeMounted,dirs[i])
		}
	} // 找到挂载目录
	for i:=0;i<len(toBeMounted);i++{
		// 注：假设目录存在不会新建
		os.MkdirAll(toBeMounted[i],664)
		commands = commands + "/bin/mount -o bind " + toBeMounted[i] + " " + serverDataDir + toBeMounted[i] + ";"
	} // 新建那些目录
	commands = commands[:len(commands) - 2]
	return commands
}