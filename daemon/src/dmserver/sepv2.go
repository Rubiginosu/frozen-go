package dmserver

import (
	"os"
	"os/exec"
	"strconv"
	"fmt"
	"os/user"
)

// 这个文件是serverEnvPrepare的第二版
// 思路学习multicraft


func (ser *ServerLocal) EnvPrepare2(){
	os.Mkdir("../servers",666)
	os.Mkdir("../servers/server" + strconv.Itoa(ser.ID),660) // 600防止其余用户读权限
	CurrentUser,_ := user.Current()
	CurrentUserUid,_ := strconv.Atoi(CurrentUser.Uid)
	os.Chown("../servers/server" + strconv.Itoa(ser.ID),CurrentUserUid,0)
	cmd := exec.Command("/usr/sbin/useradd","fg" + strconv.Itoa(ser.ID))
	err := cmd.Run()
	if err != nil {
		fmt.Println("[ERROR]" + err.Error())
		return
	}
	userCreated,_ := user.Lookup("fg" + strconv.Itoa(ser.ID))

	// TODO 完成逻辑
}