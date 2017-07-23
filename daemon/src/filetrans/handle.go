package filetrans

import (
	"net"
	"auth"
	"time"
	"os"
	"encoding/json"
)

const (
	COMMAND_AUTH = "AUTH"
	COMMAND_LIST = "LIST"
)

func (c *Command) handleCommand(conn net.Conn) {
	switch c.Command {
	case COMMAND_LIST:
		dir, err := os.Open(".")
		files,_ := dir.Readdir(-1)
		if err != nil {
			panic(err)
		}
		var localFiles []localServerFile
		for i := 0; i < len(files); i++ {
			localFiles = append(localFiles,parseFileInfoToLocalFile(files[i]))
		}
		b,_ := json.Marshal(localFiles)
		sendMessage(conn,string(b))
	}
}

func (c *Command) authCommand() bool {
	src := []byte(config.FileTransportServer.VerifyCode)
	dst := []byte(c.Args)
	return auth.Auth(src, dst)
}

func handleConnection(c net.Conn) {
	defer c.Close()
	sendMessage(c, SERVER_WELCOME)
	for {
		command := parseCommandArg(getMessage(c))
		if command.Command != COMMAND_AUTH {
			sendMessage(c, SERvER_PLEASE_AUTH)
			continue // 命令都不是AUTH的直接下一次吧。。
		} else if command.authCommand() {
			break // 验证成功就跳出循环
		} else {
			// 告诉某人认证失败。
			sendMessage(c, SERVER_AUTH_FAILED)
		}
		time.Sleep(1 * time.Second) // 防止爆破
	}
}
