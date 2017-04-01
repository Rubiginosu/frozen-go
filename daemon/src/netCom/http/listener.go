/*
 http package
 本包包含了对于http 的支持，panel和daemon的基本通信以http实现，
 原因是http可以使我们的后端工程师编写更加方便~
 Org Rubiginosu
 Team　Freeze
 Author Axoford12
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */
package http

import (
	"fmt"
	"net/http"
	"strconv" // 加载配置文件库
)

/*
 Listener 包含了对于http的监听器 基本处理器
 */
func Listen(port int){

	err := http.ListenAndServe(":" + strconv.Itoa(port),nil)
	if err != nil {
		fmt.Println("Listen and Serve error: ",err)
	}
}