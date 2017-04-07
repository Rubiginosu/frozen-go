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
	"fmt"
	"error"
	"note"
)

const DEFAULT_SECTION string = "FrozenGo"
const SECTION_TRANSLATIONS string = "translations"

// 用于返回一个map集合表示语言
func LoaderLang(languagePath,name string) {
	languagePath = "../" + languagePath
	directory, _ := os.Open(languagePath)
	fmt.Println(languagePath)
	files, err := directory.Readdir(0)
	error.ProcErr(err,"")
	for _, file := range files {
		if file.Name() == (name + ".ini") {
			langTranslations := read(languagePath + file.Name())
			languages[name] = langTranslations
		} else {
			note.Display(note.TYPE_ERROR,"Cannot get file:" + languagePath + file.Name() + ".ini")
		}
	}
}

func read(file string) (lang map[string]string){
	config := conf.SetConfig(file)
	if config.GetValue(DEFAULT_SECTION,"type") != "langPack" {
		return nil
	} else {
		langTranslations := config.ReadList()[2][SECTION_TRANSLATIONS] // 获取翻译的Map集合
		return langTranslations
	}
}