package dmserver

func IsServerAvaible(serverid int) bool {
	return serverid < len(serverSaved)-1
}
