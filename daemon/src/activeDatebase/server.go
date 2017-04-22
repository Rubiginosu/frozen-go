package activeDatebase

import (
	"gopkg.in/mgo.v2"
	"errProc"
)

func listServers(database mgo.Database) []server{
	collection := database.C(COLLECTION_SERER)
	var result []server
	err := collection.Find(nil).All(&result)
	error.ProcErr(err,"Cannot get server",nil)
	return result
}

func insertServer(database *mgo.Database,server *server) {
	collection := database.C(COLLECTION_SERER)
	err := collection.Insert(server)
	errProc.ProcErr(err,"Cannot insert server to collection..",nil)
}
