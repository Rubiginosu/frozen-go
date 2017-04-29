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



/*
一个关于中心调度器的包,用于控制,调度每个模块的正常运行.
*/
package scheduler

import (
	"manager"
	"server"
)


var serverManager *manager.ServerManager
var scheduler CentralScheduler

type CentralScheduler struct {

}

func StartScheduler() *CentralScheduler{
	startServer()
	scheduler = CentralScheduler{}
	return &scheduler
}
func (*CentralScheduler)GetServerManager() *manager.ServerManager {
	return serverManager
}
func startServer(){
	server.StartServer()
}
func getScheduler() *CentralScheduler{
	return &scheduler
}