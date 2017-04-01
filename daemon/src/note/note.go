/*
Note 包封装了打印信息的方法
 */
package note

import (
	"fmt"
	"time"
)
const TYPE_LOG string = "LOG"
const TYPE_DEBUG string = "DEBUG"
const TYPE_NOTICE string = "NOTICE"
const TYPE_ERROR string = "ERROR"

/*
Display 封装了一个打印信息的方法
 */
func Display(typeOf string,content string){
	timestamp := time.Now().Unix()
	timeUnix := time.Unix(timestamp,0)
	strTime := timeUnix.Format("2006-01-02 03:04:05 PM")
	fmt.Println("FrozenGo : " + strTime + "[ " + typeOf + " ] " + content)
}