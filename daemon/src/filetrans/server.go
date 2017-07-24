package filetrans

import (
	"conf"
	"fmt"
	"net"
	"strconv"
)
var config conf.Config
const (
	SERVER_WELCOME = "001 Frozen Go Server OK."
	SERVER_AUTH_SUCCEEDED = "002 FrozenGo Auth succeed"
	SERvER_PLEASE_AUTH = "301 Please Auth at first."
	SERVER_AUTH_FAILED = "402 Server Auth Failed"
)

func ListenAndServe(conf conf.Config){
	config = conf
	ln, err := net.Listen("tcp",":" + strconv.Itoa(conf.FileTransportServer.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c,err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}
}
