package serverManager

import (
	"conf"
	"encoding/json"
	"io/ioutil"
	"os"
)

var s []Server
/*
Command : List / Start / getStatus /
 */
func ManagerStart(ch chan string) {
	ch <- "OK"
	initial()
	for {
		handleCommand(<-ch, ch)
	}

}

type Server struct {
	ID     int
	Name   string
	Status int
}

type ServerInfomation struct {
	Name string
}

// 初始化程序
func initial() {
	config, _ := conf.GetConfig("../conf/fg.json")
	//servers := config.Smc.Servers
	// 打开文件

	b, _ := ioutil.ReadFile(config.Smc.Servers)
	json.Unmarshal(b, &s)
}

// 命令处理器
func handleCommand(command string, ch chan string) {
	switch command {

	case "List":
		outputListOfServers(ch)
	case "Create":
		s = append(s, Server{len(s), <-ch,0})
	}
}

func outputListOfServers(ch chan string) {
	b, _ := json.Marshal(s)
	ch <- string(b[:])
}

func saveServersInfo(){
	config, _ := conf.GetConfig("../conf/fg.json")
	file,_ = os.Create(config.Smc.Servers)
	// TODO Complete it.
}