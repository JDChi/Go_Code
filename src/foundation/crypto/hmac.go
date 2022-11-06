package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// hmac 就是发送方和接收方都对消息做相同步骤的 hmac 运算，然后比较计算出来的两个值是否相等
func main() {
	key := []byte("000102030405060708090a0b0c0d0e0f0")
	ciphertext := []byte("hello")
	hmacMessage := hmacSign(ciphertext, key)
	isVerified := hmacVerify(ciphertext, hmacMessage, key)
	fmt.Printf("the result is %t\n", isVerified)

}

// hmac 计算
// message 是原文信息 key 是密钥
func hmacSign(message []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	//expectedMac := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	//fmt.Printf("hello %s", expectedMac)
	return mac.Sum(nil)
}

// hmac 校验，接收方就是做同样的计算，然后跟要校验的 hmac 值做比较
func hmacVerify(message []byte, expectedMac []byte, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return hmac.Equal(mac.Sum(nil), expectedMac)
}
