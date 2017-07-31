package dmserver

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"fmt"
)

// 服务器状态码
// 已经关闭
const SERVER_STATUS_CLOSED = 0
const SERVER_STATUS_RUNNING = 1

// 按照错误码准备环境
func (server *ServerLocal) EnvPrepare() error {
	execConf,err0 := getExec(server.Executable)
	if err0 != nil {
		return err0
	}

	serverDataDir := "../servers/server" + strconv.Itoa(server.ID) // 在一开头就把serverDir算好，增加代码重用
	if _,err0 := os.Stat(serverDataDir +".loop");err0 != nil {
		fmt.Println("No loop file found!")
		//  新增 loop
		cmd := exec.Command("/bin/dd","if=/dev/zero","bs=1024",
			"count="+strconv.Itoa(execConf.MaxHardDisk),"of=../servers/server" + strconv.Itoa(server.ID) + ".loop")
		fmt.Print("Writing file...")
		err := cmd.Run()
		if err != nil {
			return err
		}
		fmt.Println("Done")

		fmt.Println("Formatting...")
		cmd2 := exec.Command("/sbin/mkfs.ext4",serverDataDir + ".loop")
		err2 := cmd2.Run()
		fmt.Println("Done")
		if err2 != nil {
			return err2
		}
		isMounted,err3 := IsDirMounted(serverDataDir)
		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
		if !isMounted {
			cmd := exec.Command("/bin/mount","-o","loop",serverDataDir + ".loop",serverDataDir)
			err := cmd.Run()
			if err != nil {
				return err
			}
		}


	}
	cmd :=  exec.Command("/bin/mount","-o","loop",serverDataDir + ".loop",serverDataDir)
	cmd.Run() // 此处不检测报错，因为有可能已经被挂载..这个貌似无法判断，干脆都挂载一次


	fileInfo, err := os.Stat(serverDataDir)
	if err != nil {
		err := server.prepareDir(serverDataDir)
		if err != nil {
			return err
		}
	} else if fileInfo.Mode() != 0660 {
		os.Chmod(serverDataDir, 0660)
	}
	os.Chown(serverDataDir,server.UserUid,server.UserUid)
	return nil
}
func (server *ServerLocal) prepareDir(serverDataDir string) error {
	err := os.MkdirAll(serverDataDir, 660)
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

func (s *ServerRun) Close() {

}

/**
name: 要寻找的exec文件
 */
func getExec(name string) (*ExecConf,error){
	file,err := os.Open("../exec/" + name + ".json") // 打开文件
	if err != nil {
		return nil,err
	}
	b,err2 := ioutil.ReadAll(file) // 读文件
	if err2 != nil {
		return nil,err
	}
	var execConf ExecConf
	err3 := json.Unmarshal(b,&execConf) // 反序列化数据
	if err3 != nil {
		return nil,err3
	}
	return &execConf,nil
}

func (server *ServerLocal) Start() error {
	server.EnvPrepare()
	execConf,err0 := getExec(server.Executable)
	if err0 != nil {
		return err0
	}
	cmd := exec.Command("server",strconv.Itoa(server.UserUid),strconv.Itoa(execConf.MaxMemory),
		"./","../servers/server" + strconv.Itoa(server.ID))

}

func outputListOfServers() Response {
	b, _ := json.Marshal(serverSaved)
	return Response{0, string(b)}
}

// 删除服务器
func (server *ServerLocal) Delete() {

	if server.Status == SERVER_STATUS_RUNNING {
		servers[server.ID].Close()
	}
	// 如果服务器仍然开启则先关闭服务器。
	// 成功关闭后，请Golang拆迁队清理违章建筑
	nowPath, _ := filepath.Abs(".")
	serverRunPath := filepath.Clean(nowPath + "/../servers/server" + strconv.Itoa(server.ID))
	os.RemoveAll(serverRunPath)
	// 清理服务器所占的储存空间
	// 违章搭建搞定以后，把这个记账本的东东也删掉
	id := searchServerByID(server.ID)
	serverSaved = append(serverSaved[:id], serverSaved[id+1:]...)
	// go这个切片是[,)左闭右开的区间，应该这么写吧~
	// 保存服务器信息。
	saveServerInfo()
}

// 搜索服务器的ID..返回index索引
// 返回-1代表没找到
func searchServerByID(id int) int {
	for i := 0; i < len(serverSaved); i++ {
		if serverSaved[i].ID == id {
			return i
		}
	}
	return -1
}
func GetServerSaved() []ServerLocal {
	return serverSaved
}

// 搜索服务器的ID..返回index索引
// 返回-1代表没找到
func searchRunningServerByID(id int) int {
	for i := 0; i < len(servers); i++ {
		if servers[i].ID == id {
			return i
		}
	}
	return -1
}
