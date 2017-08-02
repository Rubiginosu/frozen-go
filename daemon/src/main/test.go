package main

import (
	"flag"
	"syscall"
	"regexp"
	"os"
)

//#include<unistd.h>
//#include<malloc.h>
//void setug(int id){
//	while(setgid(id)!=0) sleep(1);
//  while(setuid(id)!=0) sleep(1);
//}
import "C"


func main() {
	var (
		uid     int
		mem     int64
		command string // command: ping www.baidu.com
		chroot string // eg : ./server0
	)
	flag.IntVar(&uid, "uid", 0, "uid for setuid command")
	flag.Int64Var(&mem, "mem", 1024, "mem (m) for rlimit")
	flag.StringVar(&command, "cmd", "", "Command to be run")
	flag.StringVar(&chroot,"chr","","Chroot jail for pro")
	flag.Parse()
	err := syscall.Setrlimit(syscall.RLIMIT_AS,&syscall.Rlimit{
		Cur:uint64(mem * 1048576),
		Max:uint64(mem * 1048576),
	})
	if err != nil {
		panic(err)
	}
	syscall.Chroot(chroot)
	os.Chdir(chroot + "/serverData")
	err4 := syscall.Setgroups([]int{uid})
	if err4 != nil {
		panic(err4)
	}

	C.setug(C.int(uid))

	commands := regexp.MustCompile(" +").Split(command,-1)
	err5 := syscall.Exec(commands[0], commands, []string{})
	if err5 != nil {
		panic(err5)
	}
}
