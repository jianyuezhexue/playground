package ase

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

// PKCS7Padding 填充数据以满足块大小要求
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 移除填充的数据
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, fmt.Errorf("empty data")
	}
	unpadding := int(origData[length-1])
	if unpadding > length || unpadding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	return origData[:length-unpadding], nil
}

// ECBEncrypt AES ECB模式加密
func ECBEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	encrypted := make([]byte, len(origData))

	for bs, be := 0, blockSize; bs < len(origData); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(encrypted[bs:be], origData[bs:be])
	}

	return encrypted, nil
}

// ECBDecrypt AES ECB模式解密
func ECBDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(encrypted)%blockSize != 0 {
		return nil, fmt.Errorf("encrypted data is not a multiple of the block size")
	}

	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, blockSize; bs < len(encrypted); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}
