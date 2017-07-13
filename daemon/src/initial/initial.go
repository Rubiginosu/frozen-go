/*
Author Axoford12
Team Freeze
Org Rubiginosu
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */

/*
本包包含了一些初始化的方法，用于读取配置文件 / 语言包等信息。
 */
package main

import (
	"conf"
	"fmt"
	"time"
	"os"
	"serverManager"
	"encoding/json"
)

const VERSION string = "v0.0"
const FILE_CONFIGURATION string = "../conf/fg.json"
// 用于执行一个初始化操作

func main(){
	serverManagerChan := make(chan string)
	if !(len(os.Args) > 1 && os.Args[1] == "-jump"){
		printInfo()
	}
	if !pathExists(FILE_CONFIGURATION)	{
		conf.GenerateConfig(FILE_CONFIGURATION)
	}
	config,_ := conf.GetConfig(FILE_CONFIGURATION)
	go serverManager.ManagerStart(serverManagerChan)
	if <-serverManagerChan == "OK"{
		fmt.Println("ServerManager Thread has started")
	}
	b,_ := json.Marshal(config)
	serverManagerChan <- string(b)
	//serverManagerChan <- "Create"
	//serverManagerChan <- "AVS"
}

func saveServersInfo(ch chan string){
	ch <- "List"
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
	for _,v := range []byte("Axoford12"){
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
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}