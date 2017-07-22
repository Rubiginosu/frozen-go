// 本包提供了FrozenGo专属的FTP服务器
// 包含一些奇奇怪怪的东西..
// 当然，这只是简单的FTP服务器，被动模式，TLS都并没有实现
// 以后慢慢说吧
package ftp

import (
	"fmt"
	"net"
	"strconv"
)

func ListenAndServe(port int) {
	fmt.Println("Starting FTP Server.")
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port)) // 由于是FTP服务器，使用TCP来监听
	if err != nil {
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Connection from %v established.\n", c.RemoteAddr())
		// TODO: Implement func : HandleConnection.
		go HandleConnection(c)
	}
}
