package dmserver

import (
	"os/exec"
	"strings"
)

func IsServerAvaible(serverid int) bool {
	for i := 0; i < len(serverSaved); i++ {
		if serverSaved[i].ID == serverid {
			return true
		}
	}
	return false
}

func IsDirMounted(dir string) (bool,error){
	cmd := exec.Command("/bin/df")
	resData,err := cmd.Output()
	if err != nil {
		return false,err
	}
	res := string(resData)
	n := strings.Index(res,"loop")
	return n>=0,nil
}