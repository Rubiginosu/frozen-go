/*
Error包提供了一套对于错误的处理和打印机制
 */
package error

import(
	"note"
	"os"
)
// 处理一般错误并进行打印
func ProcErr(err error,info string){
	if err != nil {
		note.Display(note.TYPE_ERROR,err.Error() + info)
		os.Exit(2)
	}

}