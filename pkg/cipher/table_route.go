package cipher

import (
	"slices"
)

type TableRouteKey struct {
	Length int
	Width  int
}

type tableRouteEncrypter struct {
	table [][]*string
	key   TableRouteKey
}

func (encrypter tableRouteEncrypter) Encrypt(to_encrypt string) string {
	if to_encrypt == "" {
		return to_encrypt
	}

	if len(to_encrypt) > encrypter.key.Length*encrypter.key.Width {
		panic("Error: key capacity doesn't cover length of message")
	}

	to_encrypt_counter_index := len(to_encrypt) - 1

	counter := 0
	to_encrypt_sliced := []rune(to_encrypt)
	for index_row := range encrypter.table {
		for index_letter := range encrypter.table[index_row] {
			letter := string(to_encrypt_sliced[counter])
			encrypter.table[index_row][index_letter] = &letter
			counter++
			if counter > to_encrypt_counter_index {
				goto END
			}

		}
	}
END:

	for index_row := range encrypter.table {
		slices.Reverse(encrypter.table[index_row])
	}

	result := ""
	for index_row := range encrypter.table {
		for _, letter := range encrypter.table[index_row] {
			if letter != nil {
				result += *letter
			}
		}
	}
	return result
}

func (decrypter tableRouteEncrypter) Decrypt(to_decrypt string) string {
	if to_decrypt == "" {
		return to_decrypt
	}

	if len(to_decrypt) > decrypter.key.Length*decrypter.key.Width {
		panic("Error: key capacity doesn't cover length of message")
	}

	to_decrypt_counter_index := len(to_decrypt) - 1

	counter := 0
	to_decrypt_sliced := []rune(to_decrypt)
	for index_row := range decrypter.table {
		for index_letter := len(decrypter.table[index_row]) - 1; index_letter >= 0; index_letter-- {
			letter := string(to_decrypt_sliced[counter])
			decrypter.table[index_row][index_letter] = &letter
			counter++
			if counter > to_decrypt_counter_index {
				goto END
			}

		}
	}
END:

	result := ""
	for index_row := range decrypter.table {
		for _, letter := range decrypter.table[index_row] {
			if letter != nil {
				result += *letter
			}
		}
	}
	return result
}
