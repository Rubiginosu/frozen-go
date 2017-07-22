package conf

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"time"
	"strconv"
	"strings"
	"math/rand"
)

type Config struct {
	ServerManagerConfig serverManagerConfig
	DaemonServerConfig  DaemonServerConfig
}

type DaemonServerConfig struct {
	Port                            int
	VerifyCode                      string
	DefaultBufLength                int
	ValidationKeyOutDateTimeSeconds float64
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
		DaemonServerConfig{52023, RandString(20), 256, 20}, // 为何选择52023？俺觉得23号这个妹纸很可爱啊
	}
	s, _ := json.MarshalIndent(v, "", "\t")
	file.Write(s)

	return nil
}

// 用于获取一个随机字符串
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}
