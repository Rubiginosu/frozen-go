package main

import (
	"fmt"
	"conf"
	"time"
	"auth"
	"dmserver"
	"os"
)

const VERSION string = "v0.2.0_Alpha"
const FILE_CONFIGURATION string = "../conf/fg.json"

var config conf.Config


/*
Command : List / Start / getStatus /
 */
func main() {
	if !(len(os.Args) > 1 && os.Args[1] == "-jump") {
		printInfo()
	}
	fmt.Println("Loading config file...")
	config, _ = conf.GetConfig(FILE_CONFIGURATION)
	fmt.Println("Config get done.")

	fmt.Println("Starting Server Manager.")
	go dmserver.StartDaemonServer(config)
	//go filetrans.ListenAndServe(config)
	fmt.Println("Starting ValidationKeyUpdater.")
	go auth.ValidationKeyUpdate(config.DaemonServer.ValidationKeyOutDateTimeSeconds)
	fmt.Println("Done,type \"?\" for help. ")
	for {
		var s string
		fmt.Scanf("%s", &s)
		processLocalCommand(s)
	}
}

func printInfo() {
	fmt.Println("  _____                        ____")
	fmt.Println("|  ___| __ ___ _______ _ __  / ___| ___")
	fmt.Println("| |_ | '__/ _\\_  / _ \\ '_ \\| |  _ / _ \\")
	fmt.Println("|  _|| | | (_) / /  __/ | | | |_| | (_) |")
	fmt.Println("|_|  |_|\\___/___\\___|_| |_|\\____|\\___/")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------")
	time.Sleep(100 * time.Microsecond)
	fmt.Print("Powered by ")
	for _, v := range []byte("Axoford12") {
		time.Sleep(240 * time.Millisecond)
		fmt.Print(string(v))
	}
	fmt.Println()
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(100 * time.Microsecond)
	fmt.Println("---------------------")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("version:" + VERSION)
	time.Sleep(1 * time.Second)
}

func processLocalCommand(c string) {
	switch c {
	case "stop":
		fmt.Println("Stopping...")
		dmserver.StopDaemonServer(config)
	case "?":
		fmt.Println("FrozenGo" + VERSION + " Help Manual -- by Axoford12")
		fmt.Println("stop: Stop the daemon.save server changes.")
		fmt.Println("status: Echo server status.")
		return
	}
}
