package serverManager

import "os"

type ServerLocal struct {
	ID         int
	Name       string
	Executable string
}

type ExecConf struct {
	Command string
	Args    []string
}

type serverIOChannel struct {
	In  chan string
	Out chan string
}

type ServerRun struct {
	ID      int
	IO      serverIOChannel
	Status  int
	InFile  *os.File
	OutFIle *os.File
}
