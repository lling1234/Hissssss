package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Encrypt(origin, key []byte) string {
	// 分组密钥
	block, _ := aes.NewCipher(key)
	// 获取密钥块长度
	size := block.BlockSize()
	// 补全码
	origin = PKCS7Padding(origin, size)
	// CBC 加密模式 (有五种模式)
	mode := cipher.NewCBCEncrypter(block, key[:size])
	encrypted := make([]byte, len(origin))
	// 加密
	mode.CryptBlocks(encrypted, origin)
	return base64.StdEncoding.EncodeToString(encrypted)
}

func Decrypt(encrypted string, key []byte) string {
	encryptedBytes, _ := base64.StdEncoding.DecodeString(encrypted)
	// 分组密钥
	block, _ := aes.NewCipher(key)
	// 密钥块长度
	size := block.BlockSize()
	// CBC 加密模式
	mode := cipher.NewCBCDecrypter(block, key[:size])
	origin := make([]byte, len(encryptedBytes))
	// 解密
	mode.CryptBlocks(origin, encryptedBytes)
	return string(PKCS7UnPadding(origin))
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	pads := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, pads...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPads := int(origData[length-1])
	return origData[:(length - unPads)]
}
