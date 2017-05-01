package main

import (
	"initial"
    "server"
	"scheduler"
)
func main(){
	initial.InitProgram()
	scheduler.StartScheduler()
}

