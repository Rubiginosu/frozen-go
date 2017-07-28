package dmserver

import (
	"auth"
	"conf"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os/exec"
	"strconv"
)

type ServerLocal struct {
	ID         int
	Name       string
	Executable string
	Status     int
}

type ExecConf struct {
	Command string
	Args    []string
}

type ServerRun struct {
	ID     int
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

/*
用于判断索引为index 的服务器是否在运行之中。
*/
func isServerRunning(index int, serverSaved []ServerLocal) bool {
	if index > len(serverSaved)-1 || index > len(servers)-1 {
		return false
	} else if serverSaved[index].Status != 1 {
		return false
	} else {
		return true
	}
}

/*
测试服务器的标准输入输出流是否可用。
*/
func ioCheck(request InterfaceRequest, c net.Conn) bool {
	// 判定OpeareID的Key是否有效
	if index := auth.FindValidationKey(request.Req.OperateID); index >= 0 {
		// 发送给User认证
		if auth.UserAuth(searchServerByID(request.Req.OperateID), request.Auth, index) {
			if isServerRunning(request.Req.OperateID, serverSaved) {
				return true
				// 所有条件满足，返回True
			} else {
				connErrorToExit("Server not running or Invalid ServerID", c)
				return false
			}
		} else {
			connErrorToExit("Auth Failed", c)
			return false
		}
	} else {
		connErrorToExit("OperateID not exist in ValidationPairs.", c)
		return false
	}
}

func StartDaemonServer(conf conf.Config) {
	config = conf
	b, _ := ioutil.ReadFile(config.ServerManager.Servers)
	json.Unmarshal(b, &serverSaved)
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(config.DaemonServer.Port)) // 默认使用tcp连接
	if err != nil {
		panic(err)
	} else {
		for {
			conn, err := ln.Accept()
			fmt.Println("New Client connected.")
			if err != nil {
				continue
			}
			go handleConnection(conn)
		}
	}

}
func StopDaemonServer() error {
	for i := 0; i < len(servers); i++ {
		servers[i].Cmd.Process.Kill()
	}
	for i := 0; i < len(serverSaved); i++ {
		serverSaved[i].Status = 0
	}
	return saveServerInfo()
}
