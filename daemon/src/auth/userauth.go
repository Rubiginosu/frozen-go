package auth

import "crypto/sha256"

func UserAuth(userServerID int, dst string, index int,pairs []ValidationKeyPairTime) bool {
	var src string
	if pairs[index].ValidationKeyPair.ID != userServerID {
		return false
	}
	src = pairs[index].ValidationKeyPair.Key
	sumSrc := sha256.Sum256([]byte(src))
	sumDst := sha256.Sum256([]byte(dst))
	if sumDst == sumSrc {
		return true
	}
	return false
}