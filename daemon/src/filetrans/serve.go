package filetrans

import (
	"net"
	"fmt"
)

func ListenAndServe(){
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

func handleConnection(c net.Conn){

}