package serverManager

import (
	"os"
	"encoding/json"
	"strconv"
	"io"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"bufio"
	"errors"
)

type ServerRuntimeConfig struct {
	Command string
	Args    []string
}

func (server *Server) Start() ([]*os.File,error) {
	server.EnvRepair()
	serverRC, err := server.loadExecutableConfig()
	if err != nil {
		return nil,errors.New("Cannot prepare server env")
	} else {
		nowPath,err := filepath.Abs(".")
		if err != nil {
			return nil,errors.New(err.Error())
		}
		serverRunPath := filepath.Clean(nowPath + "/../servers/server" + strconv.Itoa(server.ID))
		fmt.Println(serverRunPath)
		fmt.Println(serverRC)
		serverStream:= make([]*os.File,3)
		serverProcAttr := &os.ProcAttr{
			Dir:serverRunPath,
			// 下为临时标准输入输出测试
			//Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			Files:serverStream,
		}

		_,err2 := os.StartProcess("D:\\test\\echoPath.exe",[]string{""},serverProcAttr)
		if err2 != nil {
			return nil,errors.New(err2.Error())
		}
		return serverStream,nil
		// 下为读取测试
		/////////////////////////////////////////////
		//reader := bufio.NewReader(serverStream[1])
		//for {
		//	line,err := reader.ReadString('\n')
		//	if err != nil || io.EOF == err {
		//		break
		//	}
		//	fmt.Println(line) /* 此处是临时的Line操作，仅仅暂时读取出来*/
		//}
		//////////////////////////////////////////////
	}
}

// SSC: Server Self Checking
// Status CODE
// 000
// 三位二进制表示
const SSC_NO_CONFIG_FILE int = -1
const SSC_NO_SERVER_DIR = -2

// 检查运行环境
func (server *Server) selfChecking() int {
	var status int = 0
	_, err := os.Stat("../exec/" + server.Executable + ".json")
	_, err2 := os.Stat("../servers/server" + strconv.Itoa(server.ID))
	if err != nil {

		status += SSC_NO_CONFIG_FILE

	}
	if err2 != nil {
		status += SSC_NO_SERVER_DIR
	}
	return status
}

// 按照错误码修复环境
func (server *Server) EnvRepair() bool {
	statusCode := server.selfChecking()
	switch statusCode {
	case SSC_NO_SERVER_DIR:
		err := os.MkdirAll("../servers/server"+strconv.Itoa(server.ID), 0666)
		return err == nil
	case SSC_NO_CONFIG_FILE:
		defaultExec := ServerRuntimeConfig{"java", []string{"-jar","Minecraft.jar"}}
		//defaultExec := ServerRuntimeConfig{"ping",nil}
		file, err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b, err2 := json.MarshalIndent(defaultExec, "", "\t")
		io.WriteString(file, string(b)) // 写入文件
		return err == nil && err2 == nil
	case SSC_NO_CONFIG_FILE + SSC_NO_SERVER_DIR:
		// 两路一起执行
		err3 := os.MkdirAll("../servers/server"+strconv.Itoa(server.ID), 0666)
		defaultExec := ServerRuntimeConfig{"java", []string{"-jar","Minecraft.jar"}}
		//defaultExec := ServerRuntimeConfig{"ping",nil}
		file, err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b, err2 := json.MarshalIndent(defaultExec, "", "\t")
		io.WriteString(file, string(b)) // 写入文件
		return err == nil && err2 == nil && err3 == nil
	}
	return false
}
func (server *Server) loadExecutableConfig() (ServerRuntimeConfig, error) {
	var newServerRuntimeConf ServerRuntimeConfig
	b, err := ioutil.ReadFile("../exec/" + server.Executable + ".json") // 将配置文件读入
	if err != nil {
		// 若在读文件时就有异常则停止反序列化
		return newServerRuntimeConf, err
	}
	err2 := json.Unmarshal(b, &newServerRuntimeConf) //使用自带的json库对读入的东西反序列化
	if err2 != nil {
		return newServerRuntimeConf, err
	}
	return newServerRuntimeConf, nil // 返回结果
}
