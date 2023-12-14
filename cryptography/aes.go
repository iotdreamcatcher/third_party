/*
*

	@author: taco
	@Date: 2023/9/1
	@Time: 9:30

*
*/
package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// key不能泄露

func EnHex(text string) string {
	enStr := hex.EncodeToString([]byte(text))
	return enStr
}

// PKCS5填充方式
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

// Zero填充方式
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{0}, padding)

	return append(ciphertext, padtext...)
}

// PKCS5 反填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Zero反填充
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	return pkcs7UnPadding(crypted), nil
}

func AesEncryptByCTR(data, key string) (string, string) {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// 转成byte
	dataByte := []byte(data)
	keyByte := []byte(key)
	// 创建block
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	blockSize := block.BlockSize()
	// 创建偏移量iv,取秘钥前16个字符
	iv := []byte(key[:blockSize])
	// 补码
	padding := pkcs7Padding(dataByte, blockSize)
	// 加密模式
	stream := cipher.NewCTR(block, iv)
	// 定义保存结果变量
	out := make([]byte, len(padding))
	stream.XORKeyStream(out, padding)
	// 处理加密结果
	hexRes := fmt.Sprintf("%x", out)
	base64Res := base64.StdEncoding.EncodeToString(out)
	return hexRes, base64Res
}

// 解密
func AesDecryptByCTR(dataBase64, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// dataBase64转成[]byte
	decodeStringByte, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		panic(fmt.Sprintf("base64 DecodeString error: %s", err))
	}
	// 创建block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error: %s", err))
	}
	blockSize := block.BlockSize()
	// 创建偏移量iv,取秘钥前16个字符
	iv := []byte(key[:blockSize])
	// 创建Stream
	stream := cipher.NewCTR(block, iv)
	// 声明变量
	out := make([]byte, len(decodeStringByte))
	// 解密
	stream.XORKeyStream(out, decodeStringByte)
	// 解密加密结果并返回
	return string(pkcs7UnPadding(out))
}

// EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data []byte, key []byte) (string, error) {
	res, err := AesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes Aes 解密
func DecryptByAes(data string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, key)
}
