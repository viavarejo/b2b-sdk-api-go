package test

import (
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/security"
)

func TestEncrypt(t *testing.T) {
	encryptedMessage := security.Encrypt("marcos")
	if encryptedMessage == `""` {
		t.Error("Test failed, Deveria retornar a palavra criptografada")
	}
}
