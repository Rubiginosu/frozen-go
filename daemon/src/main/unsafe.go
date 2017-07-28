package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"
)

// 这里，程序是需要root权限来执行的。
// 由于需要root权限，该文件取名为Unsafe.

func main() {
	var serverId int
	flag.IntVar(&serverId, "serverId", -1, "Server id")
	flag.Parse()
	fmt.Println(serverId)
	if serverId == -1 {
		fmt.Println("E No server id input.")
		os.Exit(-1)
	}
	isOk, uid := EnvPrepare(serverId)
	if !isOk {
		fmt.Println("A fatal error occured")
	} else {
		fmt.Println(uid)
	}

}

// 这个就四原来那个准备环境啦！
// 不过这个略微有点小不同吧，哈哈！
func EnvPrepare(serverId int) (bool, int) {
	_, err := os.Stat("../servers/server" + strconv.Itoa(serverId))
	if err == nil {
		return true, -1
	}
	os.MkdirAll("../servers/server"+strconv.Itoa(serverId), 0750)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(100456)
	for {
		// 如果用户已经存在，则重新换一个数字~
		// 这个机制不会导致编写时要判定很多的这门那门
		// 但是用户多了可能会影响效率
		random = r.Intn(1004560)
		random += 100
		_, err2 := user.LookupId(strconv.Itoa(random))
		if err2 != nil {
			break
		}
	}
	if err2 := simpleRunCommand(exec.Command("/usr/sbin/useradd", "fg"+strconv.Itoa(random), "-u "+strconv.Itoa(random))); err2 != nil {
		fmt.Println(err2)
		return false, -1
	}

	err3 := simpleRunCommand(exec.Command("/bin/chown", "fg"+strconv.Itoa(random)+":fg"+strconv.Itoa(random), "../servers/server"+strconv.Itoa(serverId)))
	if err3 != nil {
		fmt.Println(err3)
		return false, -1
	}
	return true, random
}

// 简单执行
func simpleRunCommand(cmd *exec.Cmd) error {
	err := cmd.Run()
	return err
}
