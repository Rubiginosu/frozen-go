package main

import (
	"dmserver"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var exec dmserver.ExecConf
	b,_:=ioutil.ReadFile("../exec/ping.json")
	json.Unmarshal(b,&exec)
	//b,_ := json.MarshalIndent(dmserver.ExecConf{
	//	"ping",
	//	"ping 127.0.0.1",
	//	"",
	//	"",
	//	"",
	//	"",
	//	[]string{"/bin"},
	//},"","\t")
	//_,err := file.Write(b)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//file.Close()
	fmt.Println(exec)
}
