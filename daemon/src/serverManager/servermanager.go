package serverManager

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strconv"
)

/*
Command : List / Start / getStatus /
 */
func ManagerStart(ch chan string) {
	ch <- "OK"

	stringConfig := <-ch
	json.Unmarshal([]byte(stringConfig), &config)
	b, _ := ioutil.ReadFile(config.Smc.Servers)
	json.Unmarshal(b, &serverSaved)
	ch <- "OK"
	for {
		handleCommand(<-ch, ch)
	}

}



// 命令处理器
func handleCommand(command string, ch chan string) {
	switch command {

	case "List":
		outputListOfServers(ch)
	case "Create":
		serverSaved = append(serverSaved, ServerLocal{len(serverSaved), <-ch, ""})
		serverSaved[len(serverSaved) - 1].EnvRepair()
		b, _ := json.MarshalIndent(serverSaved,"","\t")
		fmt.Println(config.Smc.Servers)
		ioutil.WriteFile(config.Smc.Servers, b, 0666)
	case "Start":
		serverNameID := <- ch
		ID,err := strconv.Atoi(serverNameID)// 若传入的是整数，则比较ID
		var startingId int
		if err == nil{
			for i:=0;i<len(serverSaved);i++ {
				if serverSaved[i].ID == ID {
					startingId = i

				}
			}
		} else {
			// 不是整数判断名称
			for i:=0;i < len(serverSaved);i++{
				if  serverSaved[i].Name == serverNameID{
					startingId = i
				}
			}

		}
		stream,err := serverSaved[startingId].Start()
		if err != nil {
			panic(err)
		} else {
			serverStream = append(serverStream,stream)
		}

	}
}

func outputListOfServers(ch chan string) {
	b, _ := json.Marshal(serverSaved)
	ch <- string(b[:])
}
