package dmserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
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
	fileInfo, err := os.Stat(serverDataDir)
	if err != nil {
		err := server.prepareDir(serverDataDir)
		if err != nil {
			return err
		}
	} else if fileInfo.Mode() != 0660 {
		os.Chmod(serverDataDir, 0660)
	}
	_, err2 := user.LookupId(strconv.Itoa(userUid))
	if err2 != nil {
		cmd := exec.Command("/usr/sbin/useradd", "fg"+strconv.Itoa(server.ID), "-u "+strconv.Itoa(userUid))
		err := cmd.Run()
		return err
	}
	userNow, _ := user.Current()
	gid, _ := strconv.Atoi(userNow.Gid)
	os.Chown(serverDataDir, userUid, gid)

	server.UserUid = userUid
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

func (server *ServerLocal) Start() error {

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
