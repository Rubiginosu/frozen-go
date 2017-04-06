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

package message

import (
	"regexp"
	"strings"
)
const LANG_CN string = "chinese"

const DEFAULT_REPLACING string = "%FreezeDefault%"
var languages map[string]map[string]string = map[string]map[string]string{

}
var messages []string = []string{
	"FrozenGo daemon starting",
}
// 获取一个消息
func GetMessage(language ,log string,replacing map[string]string) string{

	langPackage := languages[language]

	if log,ok := langPackage[log]; ok {
		return processReplacing(log,replacing)
	} else {
		return log
	}
}
// 处理一次占位符替换
func processReplacing(log string,replacing map[string]string) string{
	reg,_ := regexp.Compile("rc.+?rc")
	for _,content := range reg.FindAllString(log,-1){
		replaceString := strings.Trim(content,"rc")// 去掉RC标记
		if replacer,ok := replacing[replaceString]; ok && replaceString != DEFAULT_REPLACING{
			log = strings.Replace(log,content,replacer,-1)// 不为默认值且有替换值则进行替换
		} else {
			log = strings.Replace(log,content,replaceString,-1)// 没有默认值或没替换值去掉rc即可
		}
	}
	return log
}