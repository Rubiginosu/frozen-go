package serverManager

import (
	"conf"
	"os"
)

var serverSaved []ServerLocal
var config conf.Config
var serverStream [][]*os.File
