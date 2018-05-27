package crypto

import (
	"testing"
	"bytes"
)

func TestBase64(t *testing.T) {
	msg := []byte("hello word")
	res, err := Base64Decode(Base64Encode(msg))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(msg, res) {
		t.Fatal("base64 msg decode failed")
	}
}
