package dmserver

import (
	"conf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"os"
)

func StartDaemonServer(conf conf.Config) {
	config = conf
	b, _ := ioutil.ReadFile(config.ServerManager.Servers)
	err2 := json.Unmarshal(b, &serverSaved)
	if err2 != nil{
		fmt.Println(err2)
		os.Exit(-2)
	}
	b, _ = ioutil.ReadFile(config.ServerManager.Modules)
	err3 := json.Unmarshal(b,&modules)
	if err3 != nil {
		fmt.Println(err3)
		os.Exit(-2)
	}
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(config.DaemonServer.Port)) // 默认使用tcp连接
	if err != nil {
		panic(err)
	} else {
		for {
			conn, err := ln.Accept()
			fmt.Println("[Daemon]New Client Request send.From " + conn.LocalAddr().String())
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
