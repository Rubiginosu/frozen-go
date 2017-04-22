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

/*
本包用于对于数据库的操作
FrozenGo 使用 MongoDB数据库
 */
package activeDatebase



import(
	"gopkg.in/mgo.v2"
	"errProc"
)

// 数据库链接，用于连接到MongoDB数据库
func databaseConnector(connInfo *connectionInfo) *mgo.Database{
	session,err := mgo.Dial(connInfo.DatabaseUrl)
	errProc.ProcErr(err,"Cannot connect to database:%databaseUrl%",map[string]string{
		"databaseUrl" : connInfo.DatabaseUrl,
	})
	defer session.Close()
	database := session.DB(connInfo.DatabaseName)
	return database
}
