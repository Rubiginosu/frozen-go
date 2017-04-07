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
var languages map[string]map[string]string = map[string]map[string]string{

}
// 获取一个消息
func GetMessage(languageName,log string,replacing map[string]string) string{
	if logging,ok := languages[languageName][log]; ok { // 语言包找到
		return processReplacing(logging,replacing)
	} else { // 语言包没找到
		return processReplacing(log,replacing)
	}
}
// 处理一次占位符替换
func processReplacing(log string,replacing map[string]string) string{
	reg,_ := regexp.Compile("%.+?%")
	for _,content := range reg.FindAllString(log,-1){
		replaceString := strings.Trim(content,"%")// 去掉RC标记
		if replacer,ok := replacing[replaceString]; ok{
			log = strings.Replace(log,content,replacer,-1)// 不为默认值且有替换值则进行替换
		} else {
			log = strings.Replace(log,content,replaceString,-1)// 没有默认值或没替换值去掉rc即可
		}
	}
	return log
}