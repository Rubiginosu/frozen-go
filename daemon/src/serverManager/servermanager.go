package serverManager

import (
	"conf"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

var s []Server
var config conf.Config
/*
Command : List / Start / getStatus /
 */
func ManagerStart(ch chan string) {
	ch <- "OK"

	stringConfig := <-ch
	json.Unmarshal([]byte(stringConfig), &config)
	b, _ := ioutil.ReadFile(config.Smc.Servers)
	json.Unmarshal(b, &s)
	for {
		handleCommand(<-ch, ch)
	}

}

type Server struct {
	ID         int
	Name       string
	Status     int
	Executable string
}

type ServerInfomation struct {
	Name string
}

// 命令处理器
func handleCommand(command string, ch chan string) {
	switch command {

	case "List":
		outputListOfServers(ch)
	case "Create":
		s = append(s,Server{len(s),<-ch,0,""})
		b, _ := json.MarshalIndent(s,"","\t")
		fmt.Println(config.Smc.Servers)
		ioutil.WriteFile(config.Smc.Servers, b, 0666)
	}
}

func outputListOfServers(ch chan string) {
	b, _ := json.Marshal(s)
	ch <- string(b[:])
}
