package conf

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Config struct {
	Smc  serverManagerConfig
	Lang languageConfig
}

type serverManagerConfig struct {
	Servers string
}

type languageConfig struct {
	langPath string
}

func GetConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	var v Config
	b, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return Config{}, err2
	}
	json.Unmarshal(b, &v)
	return v,nil
}


func GenerateConfig(filepath string) error{
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	var v Config = Config{
		serverManagerConfig{"data/servers"},
		languageConfig{"spk/lang/chinese.ini"},
	}
	s,_ := json.Marshal(v)
	file.Write(s)

	return nil
}
