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
func main(){
	if len(os.Args) <= 1 || os.Args[1] != "-daemon"{
		fmt.Print("Do not open me in user cmd.")
		os.Exit(-1)// 用一个参数来防止用户蜜汁点开
	}
	fmt.Println("Server Manager Has been started") // 通知主逻辑线程
	initial()
	var s string
	for{
		fmt.Scanf("%s",&s)
	}

}

type Server struct{
	ID int
	Name string
}
// 初始化程序
func initial(){
	conf := conf.SetConfig("../cnf/frozengo.ini")
	fmt.Print(conf.ReadList())
	config := conf.GetValue("ServerManager","Config")
	fmt.Print(config)
	//file,err := os.Open(config)
	//if err != nil {
	//	generateNewServerConf(config)
	//	return
	//}
	//// 打开文件
	//dec := gob.NewDecoder(file)
	//err2 := dec.Decode(&s)
	//errProc.ProcErr(err2,"Cannot decode server file",nil)

}


// 命令处理器
func handleCommand(command string){

}
func generateNewServerConf(configPath string){
	os.Create(configPath)
	file,_ := os.Open(configPath)
	enc := gob.NewEncoder(file)
	enc.Encode(s)
}