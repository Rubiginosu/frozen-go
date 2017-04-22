/*
Error包提供了一套对于错误的处理和打印机制
 */
package errProc

import(
	"os"
	"note"
	"time"
)
// 处理一般错误并进行打印
func ProcErr(err error,info string,replacing map[string]string){
	if err != nil {
		note.Display(note.TYPE_ERROR,info,replacing)
		time.Sleep(5 * time.Second)
		os.Exit(2)

	}

}
