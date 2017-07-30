package main

import (
	"dmserver"
	"encoding/json"
	"fmt"
)

func main() {
	b, _ := json.MarshalIndent(dmserver.ExecInstallConfig{
		[]dmserver.Module{
			{
				Name:     "php7",
				Download: "www.php.net",
				Chmod:    "755,root",
			},
		},
		true,
		1234,
		"http://www.baidu.com",
		dmserver.ExecConf{
			"pm",
			"php",
			[]string{"pm.phar"},
			"",
			"",
			"",
			1024,
			"stop",
		},
		"OK",
	}, "", "\t")
	fmt.Println(string(b))
}
