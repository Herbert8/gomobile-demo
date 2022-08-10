package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//初始化向量
var ivString = "4WPAd5MZ2icVRaY%"

// SetIVString 设置 IV
func SetIVString(ivStr string) {
	if len([]byte(ivStr)) == 16 {
		ivString = ivStr
	}
}

// GetIVString 获取 IV
func GetIVString() string {
	return ivString
}

// AesEncrypt AES加密
func AesEncrypt(originalData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	originalData = pkcs5Padding(originalData, blockSize)
	//originalData = pkcs7Padding(originalData, blockSize)
	iv := []byte(ivString)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherData := make([]byte, len(originalData))
	blockMode.CryptBlocks(cipherData, originalData)

	return cipherData
}

// AesDecrypt AES解密
func AesDecrypt(secretData, key []byte) []byte {
	ciphertext := secretData
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic("invalid decrypt key")
	}

	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		panic("ciphertext too short")
	}

	iv := []byte(ivString)
	if len(ciphertext)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = pkcs5UnPadding(plaintext)

	return plaintext
}

// pkcs5Padding 明文补码算法
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// pkcs5UnPadding 明文减码算法
func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

/*
  CBC加密 按照golang标准库的例子代码
  不过里面没有填充的部分,所以补上
  https://studygolang.com/articles/14251
*/

// pkcs7Padding 使用PKCS7进行填充，IOS也是7
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
