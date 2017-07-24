package filetrans

import (
	"net"
	"os"
	"encoding/json"
	"io"
	"time"
)

const (
	COMMAND_AUTH = "AUTH"
	COMMAND_LIST = "LIST"
	COMMAND_UPLOAD = "UPLO"
	COMMAND_DOWNLOAD = "DOLO"

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
	case COMMAND_UPLOAD:
		file,err := os.Create(c.Args)
		if err != nil {
			panic(err)
		}
		buf := make([]byte,10)
		for {

			_,err := conn.Read(buf)
			if err != nil || err == io.EOF{
				break
			}
			file.Write(buf)
		}
		file.Close()
	case COMMAND_DOWNLOAD:

	}
}
//
func (c *Command) authCommand() bool {
	//src := []byte(config.FileTransportServer.VerifyCode)
	//dst := []byte(c.Args)
	//return auth.Auth(src, dst)
	return true
	// TODO 实现该方法。
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
			sendMessage(c,SERVER_AUTH_SUCCEEDED)
			break // 验证成功就跳出循环
		} else {
			// 告诉某人认证失败。
			sendMessage(c, SERVER_AUTH_FAILED)
			b,_ := json.Marshal(command)
			sendMessage(c,string(b))
		}
		time.Sleep(1 * time.Second) // 防止爆破
	}
	for {
		command := parseCommandArg(getMessage(c))
		command.handleCommand(c)
	}
}
