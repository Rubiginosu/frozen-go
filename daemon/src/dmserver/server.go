package dmserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"os/user"
	"syscall"
)


// 服务器状态码
// 已经关闭
const SERVER_STATUS_CLOSED = 0
const SERVER_STATUS_RUNNING = 1

// 按照错误码准备环境
func (server *ServerLocal) EnvPrepare() error {
	userUid := server.ID + config.DaemonServer.UserIdOffset
	serverDataDir := "../servers/server" + strconv.Itoa(server.ID) // 在一开头就把serverDir算好，增加代码重用
	fileInfo,err := os.Stat(serverDataDir)
	if err  != nil {
		err := server.prepareDir(serverDataDir)
		if err != nil {
			return err
		}
	} else if fileInfo.Mode() != 0660{
		os.Chmod(serverDataDir,0660)
	}
	_,err2 := user.LookupId(strconv.Itoa(userUid))
	if err2 != nil {
		cmd := exec.Command("/usr/sbin/useradd","-s","/sbin/nologin","fg"+strconv.Itoa(server.ID),"-u " + strconv.Itoa(userUid))
		err := cmd.Run()
		return err
	}
	userNow,_ := user.Current()
	gid,_ :=strconv.Atoi(userNow.Gid)
	os.Chown(serverDataDir,userUid,gid)

	server.UserUid = userUid
	return nil
}

func (server *ServerLocal) prepareDir(serverDataDir string) error{
	err := os.MkdirAll(serverDataDir,660)
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
	s.Cmd.Process.Release()
	s.Cmd.Process.Kill()
	serverSaved[searchServerByID(s.ID)].Status = SERVER_STATUS_CLOSED
}

func (server *ServerLocal) Start() error {
	if server.Status == 1 {
		return errors.New("Server already started.")
	}
	err := server.EnvPrepare()
	if err != nil {
		return err
	}
	serverRC, err2 := server.loadExecutableConfig()
	if err2 != nil {
		// 环境准备失败
		return errors.New("Cannot prepare server env(exec file not found!")
	} else {
		// 如果Command就是一个绝对路径，不再寻找。
		execPath := serverRC.Command
		if !filepath.IsAbs(serverRC.Command) {
			var isNoFound error
			execPath, isNoFound = exec.LookPath(serverRC.Command)
			if isNoFound != nil {
				return isNoFound // 没找到抛err
			}
		}
		// 根据提供的EXEC名，搜寻绝对目录
		nowPath, err := filepath.Abs(".")
		if err != nil {
			return errors.New(err.Error())
		}
		// 取得服务器目录
		serverRunPath := filepath.Clean(nowPath + "/../servers/server" + strconv.Itoa(server.ID))
		cmd := exec.Command(execPath, serverRC.Args...)
		cmd.Dir = serverRunPath
		stdout, err := cmd.StdoutPipe()
		stdin, err2 := cmd.StdinPipe()
		if err2 != nil {
			panic(err2)
		}
		if err != nil {
			panic(err)
		}
		userNow,_ := user.Current()
		gid,_ :=strconv.Atoi(userNow.Gid)
		cmd.SysProcAttr = &syscall.SysProcAttr{Credential:&syscall.Credential{Uid:uint32(server.UserUid),Gid:uint32(gid)}}
		err3 := cmd.Start()
		if err3 != nil {
			panic(err3)
		}
		newRunningServer := ServerRun{
			ID:     server.ID,
			Cmd:    cmd,
			Stdout: stdout,
			Stdin:  stdin,
		}
		server.Status = SERVER_STATUS_RUNNING
		servers = append(servers, newRunningServer)
		return nil
	}
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
