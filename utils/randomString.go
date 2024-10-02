package utils

import (
	"math/rand"
	"time"
)

// 字母集合
const letters = "123456789asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP"

// RandomString 随机生成字符串的函数
func RandomString(n int) string {
	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}
