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
	"os"
	"conf"
)

const PATH_RESOURCE_PACK string = "../recourcePack/lang"
const DEFAULT_SECTION string = "Frozeno"
const SECTION_LANG_DECORATE string = "langDecorte"
const SECTION_TRAMSLATIONS string = "translations"
// 用于返回一个map集合表示语言
func LoaderLang() {

	directory, _ := os.Open(PATH_RESOURCE_PACK)
	files, _ := directory.Readdir(0)
	for _, file := range files {
		if !file.IsDir(){
			name,content := read(file)
			languages[name] = content
		}
	}
}

func read(file os.FileInfo) (name string,lang map[string]string){
	config := conf.SetConfig(file.Name())
	if config.GetValue(DEFAULT_SECTION,"type") != "lang" {
		return "",nil
	} else {
		languageName := config.GetValue(SECTION_LANG_DECORATE, "name")

		langContent := map[string]string{}
		for _,v := range messages{
			langContent[v] = config.GetValue(SECTION_TRAMSLATIONS,v)
		}
		return languageName,langContent
	}
}