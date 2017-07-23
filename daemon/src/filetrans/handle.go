package filetrans

import (
	"net"
	"auth"
)
const (
	COMMAND_AUTH = "AUTH"
	COMMAND_LIST = "LIST"
)
func (c *Command)handleCommand(conn net.Conn){
	switch c.Command{
	case COMMAND_LIST:

	}
}

func (c *Command)authCommand() bool{
	src := []byte(config.FileTransportServer.VerifyCode)
	dst := []byte(c.Args)
	return auth.Auth(src,dst)
}


func handleConnection(c net.Conn){
	defer c.Close()
	sendMessage(c,SERVER_WELCOME)
	for {
		command := parseCommandArg(getMessage(c))
		if command.Command != COMMAND_AUTH {
			sendMessage(c,SERvER_PLEASE_AUTH)
			continue // 命令都不是AUTH的直接下一次吧。。
		} else if command.authCommand(){
			break // 验证成功就跳出循环
		} else {
			// 告诉某人认证失败。
			sendMessage(c,SERVER_AUTH_FAILED)
		}
	}
}
