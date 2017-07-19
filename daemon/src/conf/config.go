package conf

import (
	"encoding/json"
	"os"
	"io/ioutil"
)

type Config struct {
	Smc serverManagerConfig
	Dsc DaemonServerConfig
}

type DaemonServerConfig struct {
	HttpPort   int
	DataPort   int
	VerifyCode string
}

type serverManagerConfig struct {
	Servers string
}

func GetConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		GenerateConfig("../conf/fg.json")
	}
	var v Config
	b, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return Config{}, err2
	}
	json.Unmarshal(b, &v)
	return v, nil
}

func GenerateConfig(filepath string) error {
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	var v Config = Config{
		serverManagerConfig{"../data/servers.json"},
		DaemonServerConfig{52023, 52025,"Test"}, // 为何选择52023？俺觉得23号这个妹纸很可爱啊
	}
	s, _ := json.MarshalIndent(v, "", "\t")
	file.Write(s)

	return nil
}
