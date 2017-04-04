/*
 _____                        ____       
|  ___| __ ___ _______ _ __  / ___| ___  
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \ 
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/ 

Axoford12
*/

package cmd

import (
    "bufio"
    "io"
    "os/exec"
    "error"
)
// 导入IO标准库
// 暂时实现实时输出
// 执行一个命令(底层)
// command: 命令 如 docker
// params: 参数(args)
func ExecCommand(command string,params []string,input chan string,output chan string) bool{
    cmd := exec.Command(command,params...)
    stdout, err := cmd.StdoutPipe()// 获取输出流
    if err != nil {
        // 返回错误信息时结束程序
        return false
    }
    sdtin,err := cmd.StdinPipe()
    error.ProcErr(err)
    cmd.Start() // 开始运行
    go procOutput(stdout,output)
    cmd.Wait()

    return true
}

// 处理输出流
func procOutput(stdout io.ReadCloser,chl chan string){
    reader := bufio.NewReader(stdout)//取得一个Reader
    for {
        line,err2 := reader.ReadString('\n')
        if err2 != nil || err2 == io.EOF {
            close(chl)// 错误或EOF关闭管道
            break
        }
        chl <- line//写入到数据管道
    }
    return
}
// 处理输入流
func procInput(stdin io.WriteCloser,chl chan []byte){
    writer := bufio.NewWriter(stdin)
    for data := range chl {
        writer.Write(data)
    }
}