package filetrans

import (
	"net"
	"fmt"
	"conf"
)
var config conf.Config
const (
	SERVER_WELCOME = "001 Frozen Go Server OK."
	SERvER_PLEASE_AUTH = "401 Please Auth at first."
	SERVER_AUTH_FAILED = "402 Server Auth Failed"


)

func ListenAndServe(conf conf.Config){
	config = conf
	ln, err := net.Listen("tcp",":52025")
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
