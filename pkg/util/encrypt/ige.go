package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
)

type aesIGE struct {
	cipher   cipher.Block
	iv1, iv2 []byte
}

func NewAesIGE(key, iv []byte) (*aesIGE, error) {
	if len(iv) != 32 {
		return nil, fmt.Errorf("IV length must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &aesIGE{
		cipher: block,
		iv1:    iv[:aes.BlockSize],
		iv2:    iv[aes.BlockSize:],
	}, nil
}

func (ige *aesIGE) Encrypt(src []byte) (dst []byte, err error) {
	if len(src)%aes.BlockSize != 0 {
		return nil, gerror.New("plaintext is not a multiple of the block size")
	}
	for i := 0; i < len(src); i += aes.BlockSize {
		copy(dst[i:], src[i:i+aes.BlockSize])
		for j := range dst[i : i+aes.BlockSize] {
			dst[i+j] ^= ige.iv1[j]
		}
		ige.cipher.Encrypt(dst[i:i+aes.BlockSize], dst[i:i+aes.BlockSize])
		for j := range dst[i : i+aes.BlockSize] {
			dst[i+j] ^= ige.iv2[j]
			ige.iv1[j] = dst[i+j]
		}
		copy(ige.iv2, src[i:i+aes.BlockSize])
	}
	return
}

func (ige *aesIGE) Decrypt(src []byte) (dst []byte, err error) {
	if len(src)%aes.BlockSize != 0 {
		return nil, gerror.New("ciphertext is not a multiple of the block size")
	}

	for i := 0; i < len(src); i += aes.BlockSize {
		copy(dst[i:], src[i:i+aes.BlockSize])
		for j := range dst[i : i+aes.BlockSize] {
			dst[i+j] ^= ige.iv2[j]
		}
		ige.cipher.Decrypt(dst[i:i+aes.BlockSize], dst[i:i+aes.BlockSize])
		for j := range dst[i : i+aes.BlockSize] {
			dst[i+j] ^= ige.iv1[j]
		}
		copy(ige.iv1, src[i:i+aes.BlockSize])
		for j := range dst[i : i+aes.BlockSize] {
			ige.iv2[j] = dst[i+j]
		}
	}
	return
}
