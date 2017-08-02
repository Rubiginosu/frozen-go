package dmserver

import (
	"io"
	"os/exec"
)

type Request struct {
	Method    string
	OperateID int
	Message   string
}

type Response struct {
	Status  int
	Message string
}

type InterfaceRequest struct {
	Auth string
	Req  Request
}

type ExecInstallConfig struct {
	Rely      []Module
	Success   bool
	Timestamp int
	Url       string
	StartConf ExecConf
	Message   string
	Md5       string
}

type ExecConf struct {
	Name                 string
	Command              string // 开服的指令
	StartServerRegexp    string // 判定服务器成功开启的正则表达式
	NewPlayerJoinRegexp  string // 判定新人加入的表达式
	PlayExitRegexp       string // 判定有人退出的表达式
	StoppedServerCommand string // 服务器软退出指令
	Mount                []string
}

type Module struct {
	Name     string
	Download string
	Chmod    string
	Md5      string
}

type ServerLocal struct {
	ID          int
	Name        string
	Executable  string
	Status      int
	MaxMemory   int // 内存限制
	MaxHardDisk int
}

type ServerRun struct {
	ID         int
	ToOutput   OutputInfo
	Cmd        *exec.Cmd
	StdinPipe  *io.WriteCloser
	StdoutPipe *io.ReadCloser
}
type OutputInfo struct {
	IsOutput bool
	To       chan []byte
}
type ServerAttrElement struct {
	// Set server.AttrName = AttrValue
	// eg server.MaxMemory = 1024
	AttrName  string // Attribute Name
	AttrValue string // Attribute Value
}
