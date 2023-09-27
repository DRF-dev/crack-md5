package internal

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMD5(str string) string {
	bytes := []byte(str)
	checksum := md5.Sum(bytes)
	return hex.EncodeToString(checksum[:])
}
