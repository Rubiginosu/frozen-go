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
)
// 导入IO标准库
// 暂时实现实时输出
// 执行一个命令(底层)
// command: 命令 如 docker
// params: 参数(args)
func ExecCommand(command string,params []string,c chan string) bool{
    cmd := exec.Command(command,params...)
    
    stdout, err := cmd.StdoutPipe()// 获取输出流
    if err != nil {
        // 返回错误信息时结束程序
        return false
    }

    cmd.Start() // 开始运行
    reader := bufio.NewReader(stdout)//获取reader

    // 循环读取输出流

    for {
        line,err2 := reader.ReadString('\n')
        if err2 != nil || err2 == io.EOF{
            break // IO EOF或错误 离开循环

        }
        c <- line // 将Line放入管道
    }
    cmd.Wait()
    return true
}
