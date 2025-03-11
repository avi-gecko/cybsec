package cipher_test

import (
	"testing"

	"github.com/avi-gecko/cybsec/pkg/cipher"
)

func TestElgamal(t *testing.T) {
	encrypter, err := cipher.Create(cipher.ElGamalKey{})

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	encrypted_string := encrypter.Encrypt("TestString")
	decrypted_string := encrypter.Decrypt(encrypted_string)
	if decrypted_string != "TestString" {
		t.Errorf("\nStrings are not equal:\nEncrypted: %s\nDecrypted: %s", encrypted_string, decrypted_string)
		t.FailNow()
	}
	t.Logf("\nStrings are equal:\nEncrypted: %s\nDecrypted: %s", encrypted_string, decrypted_string)

}


func TestElgamalEmptyString(t *testing.T) {
	encrypter, err := cipher.Create(cipher.ElGamalKey{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	encrypted_string := encrypter.Encrypt("")
	decrypted_string := encrypter.Decrypt(encrypted_string)
	if decrypted_string != "" {
		t.Errorf("\nStrings are not equal:\nEncrypted: %s\nDecrypted: %s", encrypted_string, decrypted_string)
		t.FailNow()
	}
	t.Logf("\nStrings are equal:\nEncrypted: %s\nDecrypted: %s", encrypted_string, decrypted_string)
}