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
}
type ExecConf struct {
	Name                 string
	Command              string   // 开服的指令
	Args                 []string // 后面的参数
	StartServerRegexp    string   // 判定服务器成功开启的正则表达式
	NewPlayerJoinRegexp  string   // 判定新人加入的表达式
	PlayExitRegexp       string   // 判定有人退出的表达式
	MaxMemory            int      // 内存限制
	StoppedServerCommand string   // 服务器软退出指令

}
type Module struct {
	Name     string
	Download string
	Chmod    string
}

type ServerLocal struct {
	ID         int
	Name       string
	Executable string
	Status     int
	UserUid    int
}

type ServerRun struct {
	ID     int
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}
