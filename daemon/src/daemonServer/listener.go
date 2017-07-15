package daemonServer

import (
	"net"
	"encoding/json"
	"crypto/sha256"
	"conf"
	"strconv"
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
	if err != nil {
		outch <- err.Error()
		panic(err)
	} else {
		for {
			conn, err := ln.Accept()
			if err != nil {
				// handle error
				continue
			}
			go handleConnection(conn)
		}
	}

}
func handleConnection(c net.Conn,) {
	welcomeMessage,_ := json.Marshal(Response{0,"Hello!"})
	c.Write(welcomeMessage)
	var request []byte
	c.Read(request)
	verifyCode := sha256.Sum256(request)
	localVerifyCode := sha256.Sum256([]byte(config.VerifyCode))
	if verifyCode == localVerifyCode {
		c.Write([]byte("OK"))
	} else {
		c.Write([]byte("Verify Error"))
		c.Close()
	}
}
