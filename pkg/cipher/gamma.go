package cipher

import (
	"crypto/rand"
)

type GammaKey struct{}

type gammaEncrypter struct {
	key   []byte
	Penis int
}

func (encrypter *gammaEncrypter) Encrypt(to_encrypt string) string {
	to_encrypt_bytes := []byte(to_encrypt)
	gamma_key := make([]byte, len(to_encrypt_bytes))
	_, err := rand.Read(gamma_key)

	if err != nil {
		panic(err)
	}

	encrypter.key = gamma_key

	for index := range to_encrypt_bytes {
		to_encrypt_bytes[index] ^= encrypter.key[index]
	}

	return string(to_encrypt_bytes)
}

func (decrypter *gammaEncrypter) Decrypt(to_decrypt string) string {
	decrypted_message := []byte(to_decrypt)

	for index := range decrypted_message {
		decrypted_message[index] ^= decrypter.key[index]
	}

	return string(decrypted_message)
}
