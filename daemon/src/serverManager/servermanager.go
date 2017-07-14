package serverManager

import (
	"conf"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strconv"
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


// 命令处理器
func handleCommand(command string, ch chan string) {
	switch command {

	case "List":
		outputListOfServers(ch)
	case "Create":
		s = append(s,Server{len(s),<-ch,0,""})
		s[len(s) - 1].generateDefaultConfig()
		b, _ := json.MarshalIndent(s,"","\t")
		fmt.Println(config.Smc.Servers)
		ioutil.WriteFile(config.Smc.Servers, b, 0666)
	case "Start":
		serverNameID := <- ch
		ID,err := strconv.Atoi(serverNameID)// 若传入的是整数，则比较ID
		if err == nil{
			for i:=0;i<len(s);i++ {
				if s[i].ID == ID {
					s[i].Start(ch)
				}
			}
		} else {
			// 不是整数判断名称
			for i:=0;i < len(s);i++{
				if  s[i].Name == serverNameID{
					s[i].Start(ch)
				}
			}
		}

	}
}

func outputListOfServers(ch chan string) {
	b, _ := json.Marshal(s)
	ch <- string(b[:])
}
