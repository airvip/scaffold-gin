package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// 生成 n 位随机字符串
func RandomString(n int) string {
	var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 生成随机数
func RandNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(1000) + 1000
}

// 加密函数
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
