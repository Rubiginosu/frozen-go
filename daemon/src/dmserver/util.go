package dmserver

func IsServerAvaible(serverid int) bool {
	for i:=0;i<len(serverSaved);i++{
		if serverSaved[i].ID == serverid{
			return true
		}
	}
	return false
}