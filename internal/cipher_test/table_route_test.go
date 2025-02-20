package cipher_test

import (
	"testing"

	"github.com/avi-gecko/cybsec/pkg/cipher"
)

func TestTableRoute(t *testing.T) {
	encrypter, err := cipher.Create(cipher.TableRouteKey{Length: 10, Width: 10})
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

func TestTableRouteEmptyString(t *testing.T) {
	encrypter, err := cipher.Create(cipher.TableRouteKey{Length: 10, Width: 10})
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

func TestTableRouteKeyCapacity(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\nPanic aren't executed")
			t.FailNow()
		}
		t.Log("\nPanic is executed")
	}()

	encrypter, err := cipher.Create(cipher.TableRouteKey{Length: 2, Width: 2})
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
