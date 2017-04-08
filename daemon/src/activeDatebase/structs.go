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

package activeDatebase

import (
	"gopkg.in/mgo.v2/bson"
)

type server struct {
	Id_        bson.ObjectId
	Name       string // 服务器名
	Token      string // 链接Token
	DaemonID   int    // 管理系统ID
	Created    int    // 创建时间
	Expiration int    // 到期时间
}

type connectionInfo struct {
	DatabaseName string
	DatabaseUrl string
}
