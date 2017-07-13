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
	"fmt"
	"time"
)

const VERSION string = "v0.0"
const FILE_CONFIGURATION string = "../conf/frozengo.ini"
// 用于执行一个初始化操作

func main(){
	printInfo()
	loadMessage()

}
func loadMessage(){
	languageName := conf.SetConfig(FILE_CONFIGURATION).GetValue("lang","language")
	languagePath := conf.SetConfig(FILE_CONFIGURATION).GetValue("lang","languagePath")
	message.LoaderLang(languagePath,languageName)
}



func printInfo() {
	fmt.Println("  _____                        ____")
	fmt.Println("|  ___| __ ___ _______ _ __  / ___| ___")
	fmt.Println("| |_ | '__/ _\\_  / _ \\ '_ \\| |  _ / _ \\")
	fmt.Println("|  _|| | | (_) / /  __/ | | | |_| | (_) |")
	fmt.Println("|_|  |_|\\___/___\\___|_| |_|\\____|\\___/")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------")
	time.Sleep(100 * time.Microsecond)
	fmt.Print("Powered by ")
	for _,v := range []byte("Axoford12"){
		time.Sleep(300 * time.Millisecond)
		fmt.Print(string(v))
	}
	fmt.Println()
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(100 * time.Microsecond)
	fmt.Println("---------------------")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("version:" + VERSION)
}


