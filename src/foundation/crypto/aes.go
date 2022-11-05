package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// AES 的加解密

// p7 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	//fmt.Printf("padding = %d\n", padding)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//fmt.Printf("padText = %s\n", padText)
	return append(data, padText...)
}

// p7 去除填充
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7 unPadding error")
	}

	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil

}

func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	// 创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取到加密块的大小，取决于传入的密钥
	blockSize := block.BlockSize()
	fmt.Printf("block size = %d\n", blockSize)
	// 填充
	encryptBytes := pkcs7Padding(data, blockSize)
	// 初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))

	// 这里的初始向量应该是使用随机数来生成才对，然后长度是块的大小
	iv := key[:blockSize]
	// 使用 cbc 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil

}

func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil

}

func EncryptByAes(data []byte, key []byte) (string, error) {
	res, err := AesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

func DecryptByAes(date string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(date)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, key)
}

// PwdKey 密钥，在AES 里，密钥长度应该为 128、192 或 256
// 这里使用长度 16 的字符串，一个字符串占 8 位，故代表的是 128 位的密钥
var PwdKey = []byte("1234asdf1234asdf")

func main() {
	var origin = []byte("123456")

	encrypt, _ := EncryptByAes(origin, PwdKey)
	decrypt, _ := DecryptByAes(encrypt, PwdKey)

	fmt.Printf("before encrypt: %s\n", origin)
	fmt.Printf("after encrypt: %s\n", encrypt)
	fmt.Printf("after decryptL %s\n", decrypt)

}
