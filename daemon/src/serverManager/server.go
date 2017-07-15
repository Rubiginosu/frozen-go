package serverManager

import (
	"os"
	"encoding/json"
	"strconv"
	"io"
	"io/ioutil"
	"path/filepath"
	"errors"
	"bufio"
	"os/exec"
)


func (server *ServerLocal) Start() ([]*os.File, error) {
	server.EnvPrepare()
	serverRC, err := server.loadExecutableConfig()
	if err != nil {
		// 环境准备失败
		return nil, errors.New("Cannot prepare server env")
	} else {
		// 根据提供的EXEC名，搜寻绝对目录
		execPath,isNoFound := exec.LookPath(serverRC.Command)
		if isNoFound != nil {
			return nil,isNoFound // 没找到抛err
		}
		nowPath, err := filepath.Abs(".")
		if err != nil {
			return nil, errors.New(err.Error())
		}
		// 取得服务器目录
		serverRunPath := filepath.Clean(nowPath + "/../servers/server" + strconv.Itoa(server.ID))
		serverStream := make([]*os.File, 3)
		serverProcAttr := &os.ProcAttr{
			Dir: serverRunPath+"/",
			//Files: serverStream,
			Files: []*os.File{os.Stdin,os.Stdout,os.Stderr},
		}

		_, err2 := os.StartProcess(execPath,append([]string{execPath},serverRC.Args...) , serverProcAttr)
		if err2 != nil {
			return nil, errors.New(err2.Error())
		}
		return serverStream, nil
	}
}



// 检查运行环境
func (server *ServerLocal) selfChecking() int {
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
func (server *ServerLocal) EnvPrepare() bool {
	statusCode := server.selfChecking()
	switch statusCode {
	case SSC_NO_SERVER_DIR:
		err := os.MkdirAll("../servers/server"+strconv.Itoa(server.ID), 0666)
		return err == nil
	case SSC_NO_CONFIG_FILE:
		defaultExec := ExecConf{"java", []string{"-jar", "Minecraft.jar"}}
		//defaultExec := ExecConf{"ping",nil}
		file, err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b, err2 := json.MarshalIndent(defaultExec, "", "\t")
		io.WriteString(file, string(b)) // 写入文件
		return err == nil && err2 == nil
	case SSC_NO_CONFIG_FILE + SSC_NO_SERVER_DIR:
		// 两路一起执行
		err3 := os.MkdirAll("../servers/server"+strconv.Itoa(server.ID), 0666)
		defaultExec := ExecConf{"java", []string{"-jar", "Minecraft.jar"}}
		//defaultExec := ExecConf{"ping",nil}
		file, err := os.Create("../exec/Minecraft.json")
		defer file.Close()
		b, err2 := json.MarshalIndent(defaultExec, "", "\t")
		io.WriteString(file, string(b)) // 写入文件
		return err == nil && err2 == nil && err3 == nil
	}
	return false
}
func (server *ServerLocal) loadExecutableConfig() (ExecConf, error) {
	var newServerRuntimeConf ExecConf
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
func (s *ServerRun) WriteReadServer() {
	go s.writeServer(s.IO.In)
	go s.readServer(s.IO.Out)
}

func  (s *ServerRun)writeServer(ch chan string){
	for range ch{
		str := <- ch
		io.WriteString(s.InFile,str)
	}
}
func (s *ServerRun)readServer(ch chan string){
	for range ch{
		reader := bufio.NewReader(s.OutFile)
		for{
			line,err := reader.ReadString('\n')
			if io.EOF == err || err != nil {
				ch <- line
			}
		}
	}
}

func (s *ServerRun)ConnectToServerPipe(in ,out chan string,id int){
	servers[id].IO.Out = out
	servers[id].IO.In = in
	s.WriteReadServer()
}

