package dmserver

import (
	"net"
	"auth"
	"os/exec"
	"io"
	"strconv"
	"fmt"
	"conf"
	"io/ioutil"
	"encoding/json"
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
用于判断索引为serverId 的服务器是否在运行之中。
 */
func isServerRunning(serverId int,serverSaved []ServerLocal) bool {
	if serverId > len(serverSaved)-1 || serverId > len(servers)-1 {
		return false
	} else if serverSaved[serverId].Status != 1 {
		return false
	} else {
		return true
	}
}
/*
测试服务器的标准输入输出流是否可用。
 */
func ioCheck(pairs []auth.ValidationKeyPairTime,request InterfaceRequest,c net.Conn) bool{
	// 判定OpeareID的Key是否有效
	if index := auth.FindValidationKey(pairs,request.Req.OperateID); index >= 0 {
		// 发送给User认证
		if auth.UserAuth(request.Req.OperateID, request.Auth, index,pairs) {
			if isServerRunning(request.Req.OperateID,serverSaved) {
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

func StartDaemonServer(config conf.Config,pairs []auth.ValidationKeyPairTime) {
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
			go handleConnection(conn,config,pairs)
		}
	}

}
func StopDaemonServer(config conf.Config) error{
	for i:=0;i<len(servers);i++{
		servers[i].Cmd.Process.Kill()
	}
	for i:=0;i<len(serverSaved);i++{
		serverSaved[i].Status = 0
	}
	return saveServerInfo(config)
}