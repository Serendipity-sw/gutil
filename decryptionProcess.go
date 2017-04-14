/**
AES加解密类库
create by gloomy 2017-03-29 23:32:31
*/
package gutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

/**
字符串加密
创建人:邵炜
创建时间:2016年3月18日09:50:36
输入参数: 需要加密的字符串
输出参数: 加密后字符串 错误对象
*/
func AesEncrypt(origData string, aESKEY []byte) (string, error) {
	origDataByte := []byte(origData)
	block, err := aes.NewCipher([]byte(aESKEY))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origDataByte = pKCS5Padding(origDataByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aESKEY[:blockSize])
	crypted := make([]byte, len(origDataByte))
	blockMode.CryptBlocks(crypted, origDataByte)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

/**
字符串解密
创建人:邵炜
创建时间:2016年3月18日09:56:20
输入参数: 需要解密的字符串  解密后字符串长度
输出参数: 解密后字符串  错误对象
*/
func AesDecrypt(crypted string, aESKEY []byte) (string, error) {
	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	block, err := aes.NewCipher(aESKEY)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	if len(cryptedByte)%blockSize != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, aESKEY[:blockSize])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = pKCS5UnPadding(origData)

	return string(origData), nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
