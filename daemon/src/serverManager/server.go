package serverManager

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type ServerRuntimeConfig struct {
	Command string
	Args    []string
}

func (server *Server) Start(infoCommunicateChan chan string) {
	_, err := os.Stat("../exec/" + server.Executable + ".json")
	_, err2 := os.Stat("../exec/" + server.Executable)
	if err == nil && err2 == nil {
		// 文件状态良好
		file, _ := os.Open("../exec/" + server.Executable + ".json")
		b, _ := ioutil.ReadAll(file)
		json.Unmarshal(b, &file)

	} else {
		// 文件不存在
		infoCommunicateChan <- "No Config or No Executable file"
	}
}

// SSC: Server Self Checking
// Status CODE
// 000
// 三位二进制表示
const SSC_NO_EXEC_FILE int = -1
const SSC_NO_CONFIG_FILE int = -2
const SSC_NO_SERVER_DIR = -4

// 检查运行环境
func (server *Server) SelfChecking() int {
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
func (server *Server) envRepair(statusCode int) bool{
	switch statusCode {
	case SSC_NO_SERVER_DIR:
		err := os.MkdirAll("../servers/server" + strconv.Itoa(server.ID),0666)
		return err == nil
	case SSC_NO_EXEC_FILE:
		return false
	case SSC_NO_CONFIG_FILE:
// TODO FINISH IT
	}

}
