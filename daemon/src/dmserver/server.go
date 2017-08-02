package dmserver

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// 服务器状态码
// 已经关闭
const SERVER_STATUS_CLOSED = 0
const SERVER_STATUS_RUNNING = 1

func (s *ServerRun) Close() {

}

func (server *ServerLocal) Start() error {
	server.EnvPrepare()
	execConf, err0 := server.loadExecutableConfig()
	if err0 != nil {
		return err0
	}
	server.MaxMemory = 1024
	cmd := exec.Command("./server", "-uid=" + strconv.Itoa(config.DaemonServer.UserId),
		"-mem="+strconv.Itoa(server.MaxMemory),
		"-chr="+"../servers/server"+strconv.Itoa(server.ID),
		 "-cmd="+execConf.Command)
	err := cmd.Start()
	if err != nil {
		return err
	}
	stdinPipe, _ := cmd.StdinPipe()
	stdoutPipe, _ := cmd.StdoutPipe()
	servers = append(servers, ServerRun{
		server.ID,
		cmd,
		stdinPipe,
		stdoutPipe,
		OutputInfo{false, nil},
	})

	go servers[len(servers)-1].ProcessOutput()
	return nil
}

func (serRun *ServerRun) ProcessOutput() {

	buf := bufio.NewReader(serRun.Stdout)
	for {
		if serRun.Cmd.ProcessState.Exited(){
			return
		}
		line, err := buf.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
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
