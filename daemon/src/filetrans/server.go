package filetrans

import (
	"conf"
	"fmt"
	"net"
	"strconv"
)

var config conf.Config

const (
	SERVER_WELCOME           = "001 Frozen Go Server OK."
	SERVER_AUTH_SUCCEEDED    = "002 FrozenGo Auth succeed"
	SERvER_PLEASE_AUTH       = "301 Please Auth at first."
	SERVER_AUTH_FAILED       = "400 Server Auth Failed"
	SERVER_ARGS_NEED_SPLIT   = "401 Server args need | to split."
	SERVER_ARGS_INVALID      = "402 Server args invalid"
	SERVER_SERVERID_INVALID  = "403 Server ID not available."
	SERVER_SERVER_INNO_ERROR = "500 Server internal error"
)

func ListenAndServe(conf conf.Config) {
	config = conf
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(conf.FileTransportServer.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}
}
