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
	"net/http"
	"../../conf"
	"../../note"
)

/*
 Listener 包含了对于http的监听器 基本处理器
 */
func Listen(){
	note.Display(note.TYPE_LOG,"Loading Conifg file...")
	config := conf.SetConfig("../config/frozengo.ini")

	note.Display(note.TYPE_LOG,"Loaded Config")
	port := config.GetValue("http","port")
	if port == "no value" {
		note.Display(note.TYPE_NOTICE,"No port infomation found ,use 52125")
		port =  "51215" // Love Girl.
	} else {
		note.Display(note.TYPE_DEBUG,"Port Founded , port:" + port + ".")
	}
	err := http.ListenAndServe(":" + port,nil)
	if err != nil {
		note.Display(note.TYPE_ERROR,"Cannot listen to port " + port)
	} else {
		note.Display(note.TYPE_DEBUG,"Listened port.")
	}
}