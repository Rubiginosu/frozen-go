/*
Author Axoford12
Team Freeze
Org Rubiginosu
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */

 package server
 // Server目录包含了
 // 监听器，逻辑API处理等操作
 import (
     "net/http"
     "note"
 )
 // 开启服务器
 func StartServer(){

 }
 func startPanelServ(port string){
     // 打印开启信息
     note.Display(note.LOG,"Starting panel server")
     http.ListenAndServe(port,nil) // 开始监听

 }
