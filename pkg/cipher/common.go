package cipher

import (
	"errors"
	"fmt"
)

type Encrypter interface {
	Encrypt(string) string
	Decrypt(string) string
}

func Create(key any) (Encrypter, error) {
	switch key := key.(type) {
	case TableRouteKey:
		table := make([][]*string, key.Width)
		for index := range table {
			table[index] = make([]*string, key.Length)
		}

		return tableRouteEncrypter{table: table, key: key}, nil

	case GammaKey:
		return &gammaEncrypter{}, nil
	case ElGamalKey:
		p := generatePrime()
		g := findPrimitiveRoot(p, findPrimitives(p))
		x := generateSecretKey(p)
		y := modExp(g, x, p)
		return &elGamalEncrypter{p: p, g: g, x: x, y: y}, nil
	default:
		return nil, errors.New("Key type: " + fmt.Sprint(key) + " doesn't exist")
	}
}
