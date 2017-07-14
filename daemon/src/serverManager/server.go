package serverManager

import (
	"os"
	"encoding/json"
	"strconv"
	"io"
)

type ServerRuntimeConfig struct {
	Command string
	Args    []string
}

func (server *Server) Start(infoCommunicateChan chan string) {
	server.EnvRepair()
}

// SSC: Server Self Checking
// Status CODE
// 000
// 三位二进制表示
const SSC_NO_EXEC_FILE int = -1
const SSC_NO_CONFIG_FILE int = -2
const SSC_NO_SERVER_DIR = -4

// 检查运行环境
func (server *Server) selfChecking() int {
	var status int = 0
	_, err := os.Stat("../exec/" + server.Executable + ".json")
	_, err2 := os.Stat("../exec/" + server.Executable)
	_, err3 := os.Stat("../servers/server" + strconv.Itoa(server.ID))
	if err == nil {
		// 文件状态良好
		status += SSC_NO_EXEC_FILE
	}
	if err2 == nil {
		status += SSC_NO_CONFIG_FILE
	}
	if err3 == nil {
		status += SSC_NO_SERVER_DIR
	}
	return status
}

// 按照错误码修复环境
func (server *Server) EnvRepair() bool{
	statusCode := server.selfChecking()
	switch statusCode {
	case SSC_NO_SERVER_DIR:
		err := os.MkdirAll("../servers/server" + strconv.Itoa(server.ID),0666)
		return err == nil
	case SSC_NO_EXEC_FILE:
		return false
	case SSC_NO_CONFIG_FILE:
		defaultExec := ServerRuntimeConfig{"java -jar Minecraft.jar",nil}
		//defaultExec := ServerRuntimeConfig{"ping",nil}
		file,err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b,err2 :=  json.MarshalIndent(defaultExec,"","\t")
		io.WriteString(file,string(b))// 写入文件
		return err == nil && err2 ==nil
	case SSC_NO_CONFIG_FILE + SSC_NO_SERVER_DIR:
		// 两路一起执行
		err3 := os.MkdirAll("../servers/server" + strconv.Itoa(server.ID),0666)
		defaultExec := ServerRuntimeConfig{"java -jar Minecraft.jar",nil}
		//defaultExec := ServerRuntimeConfig{"ping",nil}
		file,err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b,err2 :=  json.MarshalIndent(defaultExec,"","\t")
		io.WriteString(file,string(b))// 写入文件
		return err == nil && err2 == nil && err3 == nil
	}
	return false
}
