package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	md := md5.New()
	md.Write([]byte(value))
	return hex.EncodeToString(md.Sum(nil))
}
