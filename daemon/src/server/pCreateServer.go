/*
Author Axoford12
Team Freeze
Org Rubiginosu
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */
package server

/*
本文件主要包含了Panel的创建服务器方法
创建服务器使用公共返回值，但Data 数据有所不通
这里会新声明一个struct .来表示创建服务器的返回值
*/

import (
	"net/http"
)

// 服务器信息的描述
type ServerInfo struct {
	Port  int    // 服务器的端口
	Owner string // 所有用户(管理ID)
	Id    int    // 服务器ID
	Info string // 关于服务器的一些信息 Json格式
}

// 本方法作为参数传入http.HandleFunc
func createServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 解析参数
	var answer Answer
	var serverInfo ServerInfo
	if auth(getAuthHashCode(r)) {
		port := r.Form["port"][0] // 设置端口
		owner := r.Form["owner"][0]

	} else {
		authInvalid(&answer)
	}

}
