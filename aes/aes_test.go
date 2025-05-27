package ase

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

// origData := []byte("Hello World") // 待加密的数据
// key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
// log.Println("原文：", string(origData))

// log.Println("------------------ CBC模式 --------------------")

// encrypted := AesEncryptCBC(origData, key)
// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
// log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
// decrypted := AesDecryptCBC(encrypted, key)
// log.Println("解密结果：", string(decrypted))

// log.Println("------------------ ECB模式 --------------------")
// encrypted = AesEncryptECB(origData, key)
// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
// log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
// decrypted = AesDecryptECB(encrypted, key)
// log.Println("解密结果：", string(decrypted))

// log.Println("------------------ CFB模式 --------------------")
// encrypted = AesEncryptCFB(origData, key)
// log.Println("密文(hex)：", hex.EncodeToString(encrypted))
// log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
// decrypted = AesDecryptCFB(encrypted, key)
// log.Println("解密结果：", string(decrypted))
func TestECBEncrypt(t *testing.T) {
	origData := `{"saleOrderCode":"2022080101353300","shipOrderId":"2022000023","expressNo":"SF1387143774070","orderCreateTime":"2022-08-01 16:58:00","consigneeNo":"18792592230"}` // 待加密的数据
	key := "ZyXwV%uTsRqPoNmLkJiHgFeDcBa9Kdd5"                                                                                                                                       // 加密的密钥

	// 加密
	encrypted, err := ECBEncrypt([]byte(origData), []byte(key))
	if err != nil {
		t.Errorf("加密失败: %v", err)
	}
	expected := "RvbsPgf9tAdxpC9R99c4QnhyCoRFQgZ7r+G72M04aP8PsocUcGK8hSrN+BN0SA6M72vPjszMu8Kn90Goz2oT2UBmoTmMYZWvHU+SjfCVLTnX7tXpGn5sR8yNzKcBlf1KQ/OPNA9GVKKJLXCQQvofhioEgk3Idmr5E7Nz6LA+cf5bhV6YfBlVsOR3ZfbU4sbA80cOTJSyBkBEDLb28t1p5J4XCbDDc31WQljVTjhqIVI="
	fmt.Println("期待结果:", expected)
	fmt.Printf("加密结果: %s\n", base64.StdEncoding.EncodeToString(encrypted))
	assert.Equal(t, expected, base64.StdEncoding.EncodeToString(encrypted))
}

func TestECBDecrypt(t *testing.T) {
	origData := `{"saleOrderCode":"2022080101353300","shipOrderId":"2022000023","expressNo":"SF1387143774070","orderCreateTime":"2022-08-01 16:58:00","consigneeNo":"18792592230"}` // 待加密的数据
	key := "ZyXwV%uTsRqPoNmLkJiHgFeDcBa9Kdd5"                                                                                                                                       // 加密的密钥
	encryptedCode := "RvbsPgf9tAdxpC9R99c4QnhyCoRFQgZ7r+G72M04aP8PsocUcGK8hSrN+BN0SA6M72vPjszMu8Kn90Goz2oT2UBmoTmMYZWvHU+SjfCVLTnX7tXpGn5sR8yNzKcBlf1KQ/OPNA9GVKKJLXCQQvofhioEgk3Idmr5E7Nz6LA+cf5bhV6YfBlVsOR3ZfbU4sbA80cOTJSyBkBEDLb28t1p5J4XCbDDc31WQljVTjhqIVI="

	// 加密
	encryptedCodeBytes, _ := base64.StdEncoding.DecodeString(encryptedCode)
	encrypted, err := ECBDecrypt(encryptedCodeBytes, []byte(key))
	if err != nil {
		t.Errorf("解密失败: %v", err)
	}
	fmt.Println("期待结果:", origData)
	fmt.Printf("解密结果: %s\n", string(encrypted))
	assert.Equal(t, origData, string(encrypted))
}
