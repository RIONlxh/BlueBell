package utils

import (
	"BlueBell/config"
	"crypto/md5"
	"encoding/hex"
)

// MD5 加密
func Md5Encrypt(str string) string {
	h := md5.New()
	h.Write([]byte(config.Conf.Secret))
	return hex.EncodeToString(h.Sum([]byte(str)))
}
