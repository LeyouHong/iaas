package crypto

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestAES(t *testing.T) {
	input := []byte("hello world")
	iv := []byte("532b6195636c6127")[:aes.BlockSize]
	key := []byte("532b6195636c61279a010000")

	encrypted := make([]byte, len(input))
	err := EncryptAES(encrypted, input, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	decrypted := make([]byte, len(input))
	err = DecryptAES(decrypted, encrypted, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(decrypted, input) {
		t.Fatal("AES msg decode failed")
	}
}
