package filetrans

import (
	"net"
	"fmt"
	"io"
	"auth"
	"os"
)

func getMessage(c net.Conn) []byte{
	b := make([]byte,256)
	length,err:= c.Read(b)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b[:length-1]
}

func sendMessage(c net.Conn,message string) bool {
	_,err := io.WriteString(c,message)
	return err == nil
}

func parseCommandArg(data []byte) *Command{
	// 前四位设置为Command ,第六位到最后是Arg
	//   AAAA      BBBBBBBBBB
	// Command       Arg
	return &Command{string(data[:4]),string(data[5:])}
}

func parseFileInfoToLocalFile(f os.FileInfo) localServerFile{
	return localServerFile{
		f.Name(),
		f.Mode().String(),
		f.IsDir(),
		f.Size(),
		f.ModTime().Unix(),
	}
}