package main

import (
	"os"
	"fmt"
	"conf"
	"encoding/gob"
)

var s []Server
/*
Command : List / Start / getStatus /
 */
func main() {
	if len(os.Args) <= 1 || os.Args[1] != "-daemon" {
		fmt.Print("Do not open me in user cmd.")
		os.Exit(-1) // 用一个参数来防止用户蜜汁点开
	}
	fmt.Println("Server Manager Has been started") // 通知主逻辑线程
	initial()
	var s string
	for {
		fmt.Scanf("%s", &s)
	}

}

type Server struct {
	ID           int
	Name         string
	Status       int
	OnlinePlayer int
}

// 初始化程序
func initial() {
	config, _ := conf.GetConfig("../conf/fg.json")
	//servers := config.Smc.Servers
	// 打开文件
	file, _ := os.Open(config.Smc.Servers)
	dec := gob.NewDecoder(file)
	err2 := dec.Decode(&s)
	if err2 != nil {
		panic(err2)
	}
}

// 命令处理器
func handleCommand(command string) {

}
