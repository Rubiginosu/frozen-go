package conf

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type config struct {
	Smc  serverManagerConfig
	Lang languageConfig
}

type serverManagerConfig struct {
	Servers string
}

type languageConfig struct {
	langPath string
}

func getConfig(filename string) (config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var v config
	b, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return nil, err2
	}
	json.Unmarshal(b, &v)
	return v,nil
}


func generateConfig(filepath string) error{
	_, err := os.Create(filepath)
	if err != nil {
		return err
	}
	var v config = config{
		serverManagerConfig{"data/servers"},
		languageConfig{"spk/lang/chinese.ini"},
	}
	s,_ := json.Marshal(v)
	ioutil.WriteFile(filepath,s,0666)
}
