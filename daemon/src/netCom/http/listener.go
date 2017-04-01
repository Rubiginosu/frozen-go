/*
 http package
 本包包含了对于http 的支持，panel和daemon的基本通信以http实现，
 原因是http可以使我们的后端工程师编写更加方便~
 Org Rubiginosu
 Team　Freeze
 Author Axoford12
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */
package http

import (
	"net/http"
	"../../conf"
	"../../note"
	"fmt"
	"unsafe"
	"reflect"
	"encoding/pem"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)

/*
 Listener 包含了对于http的监听器 基本处理器
 */
func Listen(){
	note.Display(note.TYPE_LOG,"Loading Conifg file...")
	config := conf.SetConfig("../config/frozengo.ini")
	http.HandleFunc("/token.gateway",processToken)
	note.Display(note.TYPE_LOG,"Loaded Config")
	port := config.GetValue("http","port")
	if port == "no value" {
		note.Display(note.TYPE_NOTICE,"No port infomation found ,use 52125")
		port =  "51215" // Love Girl.
	} else {
		note.Display(note.TYPE_DEBUG,"Port Founded , port:" + port + ".")
	}
	err := http.ListenAndServe(":" + port,nil)
	if err != nil {
		note.Display(note.TYPE_ERROR,"Cannot listen to port " + port)
	} else {
		note.Display(note.TYPE_DEBUG,"Listened port.")
	}
}
/*
处理Panel 服务器发来的信息并生成token
 */
func processToken(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	publicKey := r.Form["public_key"]

	if publicKey == nil{
		// 没有收到 public key
		note.Display(note.TYPE_LOG,"HTTP Request Found .Cannot find publickey in POST or GET")

	} else {
		note.Display(note.TYPE_NOTICE,"HTTP Request vailed.")
		token := Krand(20)
		note.Display(note.TYPE_NOTICE,"New Client generated" + ByteToString(token))
		encrypted,err := RsaEncrypt(StringToByte(&(publicKey[0])),token)
		if err != nil {
			note.Display(note.TYPE_ERROR,"Encrypt failed")
			fmt.Fprintf(w, "Encrypt Failded")
		} else {
			fmt.Fprintf(w, ByteToString(encrypted))
		}

	}

}
// Byte数组转String
func ByteToString(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}
// String转Byte数组
func StringToByte(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

// 加密
func RsaEncrypt(publicKey,origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
