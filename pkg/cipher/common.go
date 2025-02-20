package cipher

import (
	"errors"
	"fmt"
)

type Encrypter interface {
	Encrypt(string) string
	Decrypt(string) string
}

func Create(key interface{}) (Encrypter, error) {
	switch key := key.(type) {
	case TableRouteKey:
		table := make([][]*string, key.Width)
		for index := range table {
			table[index] = make([]*string, key.Length)
		}

		return tableRouteEncrypter{table: table, key: key}, nil

	case GammaKey:
		return &gammaEncrypter{Penis: 66}, nil
	case ElGamalKey:
		return elGamalEncrypter{}, nil
	default:
		return nil, errors.New("Key type: " + fmt.Sprint(key) + " doesn't exist")
	}
}
