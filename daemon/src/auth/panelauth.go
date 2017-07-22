package auth

import "crypto/sha256"

func Auth(src,dst []byte) bool {
	verify := sha256.Sum256(dst)
	verifyAuth := sha256.Sum256(src)
	return verify == verifyAuth
}
