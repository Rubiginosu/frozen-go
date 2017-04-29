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
	"conf"
)


var serverManager *manager.ServerManager
var schedule CentralScheduler
var config *Configure
type CentralScheduler struct {
	Event EventScheduler
}

type  Configure struct {
	Config []map[string]map[string]string
}


func StartScheduler() *CentralScheduler{
	startServer()
	loadConfig()
	schedule = CentralScheduler{}
	return &schedule
}


func GetScheduler() *CentralScheduler{
	return &schedule
}

// Get ServerManager
func (c *CentralScheduler)GetServerManager() *manager.ServerManager {
	return serverManager
}

//  start Server listen to a certain port
func startServer(){
	server.StartServer()
}

//
func loadConfig(){
	config.Config = conf.SetConfig("../config/frozengo.conf").ReadList()
}


func (c *CentralScheduler)GetConfig() *Configure{
	return config
}
