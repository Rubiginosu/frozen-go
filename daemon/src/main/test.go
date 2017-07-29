package main

import (
	"net"
	"encoding/json"
	"io"
	"os"
	"fmt"
	"time"
)
type Request struct {
	Method    string
	OperateID int
	Message   string
}

type Response struct {
	Status  int
	Message string
}

type InterfaceRequest struct {
	Auth string
	Req  Request
}
type ValidationKeyPair struct {
	ID  int // 该ID对应服务器。
	Key string
}
type ValidationKeyPairTime struct {
	ValidationKeyPair ValidationKeyPair
	GeneratedTime     time.Time
}

func main(){
	c,err := net.Dial("tcp","127.0.0.1:52023")
	if err != nil {
		panic(err)
	}

	b,_ :=json.Marshal(InterfaceRequest{"uH1C0ZKkx9SaNJJXyslV",Request{"Start",5,"123"}})
	c.Write(b)
	io.Copy(os.Stdout,c)
}
func getValidationKey() string{
	d,_ := net.Dial("tcp","127.0.0.1:52023")
	b,_ :=json.Marshal(InterfaceRequest{"uH1C0ZKkx9SaNJJXyslV",Request{"GetPairs",5,"123"}})
	d.Write(b)
	res := make([]byte,1024)
	length ,_:= d.Read(res)
	fmt.Println(string(res))
	var resp Response
	var pair ValidationKeyPairTime
	json.Unmarshal(res[:length],&resp)
	json.Unmarshal([]byte(resp.Message),&pair)
	fmt.Println(pair.ValidationKeyPair.Key)
	d.Close()
	return pair.ValidationKeyPair.Key
}