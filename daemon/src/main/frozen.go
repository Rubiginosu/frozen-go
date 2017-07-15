package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"conf"
	"os"
	"time"
	"sync"
	"io"
	"net"
	"strconv"
	"bufio"
	"crypto/sha256"
	"errors"
	"os/exec"
	"path/filepath"
)
const VERSION string = "v0.0"
const FILE_CONFIGURATION string = "../conf/fg.json"
var serverSaved []ServerLocal
var config conf.Config
var servers []ServerRun
// SSC: ServerLocal Self Checking
// Status CODE
// 00
// 二进制表示
const SSC_NO_CONFIG_FILE int = -1
const SSC_NO_SERVER_DIR = -2

// 服务器状态码

// 已经关闭
const SERVER_STATUS_CLOSED = 0
// 正在运行
const SERVER_STATUS_RUNNING = 1


type ServerLocal struct {
	ID         int
	Name       string
	Executable string
	Status     int
}

type ExecConf struct {
	Command string
	Args    []string
}

type ServerRun struct {
	ID      int
	Status  int
	InFile  *os.File
	OutFile *os.File
	Proc    *os.Process
}

type Request struct {
	Method    string
	OperateID int
	Message   string
}

type Response struct {
	Status  int
	Message string
}


/*
Command : List / Start / getStatus /
 */
func main() {
	wg := sync.WaitGroup{}
	if !(len(os.Args) > 1 && os.Args[1] == "-jump") {
		printInfo()
	}
	config,_ = conf.GetConfig(FILE_CONFIGURATION)
	b, _ := ioutil.ReadFile(config.Smc.Servers)
	json.Unmarshal(b, &serverSaved)
	fmt.Println("Started Server Manager.")
	fmt.Println("Online...")
	go StartDaemonServer()
	wg.Add(1)
	wg.Wait()
}


// 命令处理器
func handleRequest(request Request) Response{
	switch request.Method {

	case "List":
		return outputListOfServers()
	case "Create":
		serverSaved = append(serverSaved, ServerLocal{len(serverSaved), request.Message, "", 0})

		serverSaved[len(serverSaved)-1].EnvPrepare()
		// 序列化b来储存。
		b, _ := json.MarshalIndent(serverSaved, "", "\t")

		// 新创建的服务器写入data文件
		ioutil.WriteFile(config.Smc.Servers, b, 0666)

		// 写入数据
		servers = append(servers, ServerRun{ID: len(servers), Status: 0, })

		return Response{
			0,
			"OK",
		}
	case "Start":
		// 操作ID
		serverSaved[request.OperateID].Start()
		return Response{
			0, "OK",
		}
	case "Stop":
		servers[request.OperateID].Close()
		return Response{
			0, "OK",
		}
	}
	return Response{
		-1,"Unexpected err",
	}
}

func outputListOfServers() Response{
	b, _ := json.Marshal(serverSaved)
	return Response{0,string(b)}
}
func printInfo() {
	fmt.Println("  _____                        ____")
	fmt.Println("|  ___| __ ___ _______ _ __  / ___| ___")
	fmt.Println("| |_ | '__/ _\\_  / _ \\ '_ \\| |  _ / _ \\")
	fmt.Println("|  _|| | | (_) / /  __/ | | | |_| | (_) |")
	fmt.Println("|_|  |_|\\___/___\\___|_| |_|\\____|\\___/")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------")
	time.Sleep(100 * time.Microsecond)
	fmt.Print("Powered by ")
	for _, v := range []byte("Axoford12") {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(string(v))
	}
	fmt.Println()
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(100 * time.Microsecond)
	fmt.Println("---------------------")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("version:" + VERSION)
}
func StartDaemonServer() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(config.Dsc.Port)) // 默认使用tcp连接
	if err != nil {
		panic(err)
	} else {
		for {
			conn, err := ln.Accept()
			fmt.Println("New Client connected.")
			if err != nil {
				continue
			}
			go handleConnection(conn)
		}
	}

}
func Auth(c net.Conn) bool {
	fmt.Println("Connector Auth...")
	var requestBytes []byte
	reader := bufio.NewReader(c)
	for{
		requestTemp, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF{
			requestBytes = requestTemp
			break
		}

	}
	var request Request
	json.Unmarshal(requestBytes,&request)
	verifyCode := sha256.Sum256([]byte(request.Message))
	localVerifyCode := sha256.Sum256([]byte(config.Dsc.VerifyCode))
	fmt.Println([]byte(config.Dsc.VerifyCode))
	if verifyCode == localVerifyCode {
		fmt.Println("Auth passed")
		return true
	} else {
		fmt.Println("Auth FAILED")
		return false
	}
}
func handleConnection(c net.Conn) {
	if Auth(c) {
		fmt.Println("Client Auth ok,process commands")
		c.Write([]byte("Auth Ok"))
		writer := bufio.NewWriter(c)
		reader := bufio.NewReader(c)
		var request Request
		for {
			var response Response
			requestBytes, err := reader.ReadBytes('\n')
			for {
				requestBytes, err = reader.ReadBytes('\n')
				if err != nil || err == io.EOF {
					break
				}
			}
			var request Request
			err2 := json.Unmarshal(requestBytes,&request)
			if err != nil {
				continue
			} else if err2 != nil{
				response.Message = err.Error()
				response.Status = -1
			} else if request.Method != "GetInput" && request.Method != "GetOutput"{
				response = handleRequest(request)
			} else {
				break
			}
			b,_ := json.Marshal(response)
			writer.Write(b)
			writer.Flush()
		}
		if request.Method == "GetOutput"{
			io.Copy(c,servers[request.OperateID].OutFile)
		} else if request.Method == "GetInput"{
			io.Copy(servers[request.OperateID].InFile,c)
		}


	} else {
		c.Close()
	}
}

func (server *ServerLocal) Start() error {
	server.EnvPrepare()
	serverRC, err := server.loadExecutableConfig()
	if err != nil {
		// 环境准备失败
		return errors.New("Cannot prepare server env")
	} else {
		// 根据提供的EXEC名，搜寻绝对目录
		execPath, isNoFound := exec.LookPath(serverRC.Command)
		if isNoFound != nil {
			return isNoFound // 没找到抛err
		}
		nowPath, err := filepath.Abs(".")
		if err != nil {
			return errors.New(err.Error())
		}
		// 取得服务器目录
		serverRunPath := filepath.Clean(nowPath + "/../servers/server" + strconv.Itoa(server.ID))
		serverStream := make([]*os.File, 3)
		serverProcAttr := &os.ProcAttr{
			Dir: serverRunPath + "/",
			Files: serverStream,
			//Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		}

		proc, err2 := os.StartProcess(execPath, append([]string{execPath}, serverRC.Args...), serverProcAttr)
		if err2 != nil {
			return errors.New(err2.Error())
		}
		newRunningServer := ServerRun{
			ID:      server.ID,
			InFile:  serverStream[0],
			OutFile: serverStream[1],
			Proc:proc,
		}
		server.Status = 1
		servers = append(servers, newRunningServer)
		return nil
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

func (s *ServerRun) Close() {
	s.Proc.Release()
	s.Proc.Kill()
	s.OutFile.Close()
	s.InFile.Close()
}

