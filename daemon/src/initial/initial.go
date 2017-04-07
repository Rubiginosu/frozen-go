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
本包包含了一些初始化的方法，用于读取配置文件 / 语言包等信息。
 */
package initial

import (
	"conf"
	"message"
)

const FILE_CONFIGURATION string = "../config/frozengo.ini"
// 用于执行一个初始化操作

func InitProgram(){
	printInfo()
	loadMessage()
}
func loadMessage(){
	languageName := conf.SetConfig(FILE_CONFIGURATION).GetValue("lang","language")
	languagePath := conf.SetConfig(FILE_CONFIGURATION).GetValue("lang","languagePath")
	message.LoaderLang(languagePath,languageName)
}