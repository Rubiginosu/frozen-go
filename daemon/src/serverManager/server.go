package serverManager

import (
	"os"
	"encoding/json"
	"io/ioutil"
)

type ServerRuntimeConfig struct{
	Command string
	String string
}

func (server *Server) Start() {
	_, err := os.Stat("../exec/" + server.Executable + ".json")
	if err == nil {
		// 文件状态良好
		file,_ := os.Open("../exec/" + server.Executable + ".json")
		b,_ := ioutil.ReadAll(file)
		json.Unmarshal(b,&file)
	}
}
