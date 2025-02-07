// Package encrypt
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/forgoer/openssl"
	"io"
)

// AesECBEncrypt 加密
func AesECBEncrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
}

// AesECBDecrypt 解密
func AesECBDecrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
}

// MustAesECBEncryptToString
// Return the encryption result directly. Panic error
func MustAesECBEncryptToString(bytCipher, key string) string {
	dst, err := AesECBEncrypt([]byte(bytCipher), []byte(key))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(dst)
}

// MustAesECBDecryptToString
// Directly return decryption result, panic error
func MustAesECBDecryptToString(bytCipher, key string) string {
	dst, err := AesECBDecrypt([]byte(bytCipher), []byte(key))
	if err != nil {
		panic(err)
	}
	return string(dst)
}

// AesCTREncrypt 加密
func AesCTREncrypt(plaintext, key []byte) (dst, iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// CTR模式的IV必须与密钥一样长
	iv = make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 返回密文和IV
	return ciphertext, iv, nil
}

// AesCTRDecrypt 解密
func AesCTRDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
