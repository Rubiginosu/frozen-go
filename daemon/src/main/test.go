package main

import(
    "fmt"
    "servMngr/cmd"
)

func main(){
    channel := make(chan string)
    go cmd.ExecCommand("ping",[]string{"127.0.0.1"},channel)
    for i:= range(channel){
        fmt.Print(i)
    }
}
