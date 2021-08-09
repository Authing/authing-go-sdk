package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/Authing/authing-go-sdk/lib/constant"
)

func RsaEncrypt(plainText string) string {
	//pem解码
	//block, _ := pem.Decode([]byte(constant.PublicKey))
	block, _ := base64.StdEncoding.DecodeString(constant.PublicKey)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		panic(err)
	}
	//返回密文
	return base64.StdEncoding.EncodeToString(cipherText)
}
