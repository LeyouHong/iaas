package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func DecryptAES(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}
