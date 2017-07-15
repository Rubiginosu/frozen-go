package daemonServer

import (
	"net"
	"encoding/json"
	"crypto/sha256"
	"conf"
	"strconv"
	"fmt"
	"bufio"
	"bytes"
)

type Response struct {
	Status  int
	Message string
}
var config conf.DaemonServerConfig
func StartDaemonServer(inch chan string,outch chan string) {
	configStr := <- inch

	json.Unmarshal([]byte(configStr),&config)
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(config.Port)) // 默认使用tcp连接
	defer ln.Close()
	if err != nil {
		outch <- err.Error()
		panic(err)
	} else {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go handleConnection(conn)

		}
	}

}
func handleConnection(c net.Conn) {
	if welcomeAuth(c) {
		c.Write([]byte("Auth Ok"))

	} else {
		c.Close()
	}
}

func welcomeAuth(c net.Conn) bool{
	w,_:= json.Marshal(Response{0,"Hello!"})
	c.Write(w)
	var request []byte

	reader := bufio.NewReader(c)
	request,_ = reader.ReadBytes('\n')
	bytes.TrimSpace(request)
	fmt.Println(request)
	verifyCode := sha256.Sum256(request)
	localVerifyCode := sha256.Sum256([]byte(config.VerifyCode))
	fmt.Println([]byte(config.VerifyCode))
	if verifyCode == localVerifyCode {
		return true
	} else {
		return false
	}
}
