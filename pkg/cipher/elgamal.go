package cipher

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type ElGamalKey struct{}

type elGamalEncrypter struct {
	y uint64
	g uint64
	p uint64
	x uint64
}


func (e *elGamalEncrypter) Encrypt(to_encrypt string) string {
	var message []string
	a := modExp(e.g, e.x, e.p)

	for _, char := range to_encrypt {
		b := (uint64(char) * modExp(e.y, e.x, e.p)) % e.p
		message = append(message,
			strconv.FormatUint(a, 10),
			strconv.FormatUint(b, 10),
		)
	}

	return strings.Join(message, " ")
}

func (e *elGamalEncrypter) Decrypt(to_decrypt string) string {
	var decrypted strings.Builder
	splitted := strings.Split(to_decrypt, " ")

	for i := 0; i < len(splitted)-1; i += 2 {
		a, _ := strconv.ParseUint(splitted[i], 10, 64)
		b, _ := strconv.ParseUint(splitted[i+1], 10, 64)

		ax := modExp(a, e.x, e.p)
		inverse, _ := modInverse(ax, e.p)
		decrypted.WriteRune(rune((b * inverse) % e.p))
	}

	return decrypted.String()
}


func generatePrime() uint64 {
	for {
		p, _ := rand.Prime(rand.Reader, 31)
		if isPrime(p.Uint64()) {
			return p.Uint64()
		}
	}
}

func isPrime(n uint64) bool {
	if n <= 1 {
		return false
	}
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimitives(p uint64) []uint64 {
	n := p - 1
	var factors []uint64
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func findPrimitiveRoot(p uint64, factors []uint64) uint64 {
	for g := uint64(2); g < p; g++ {
		ok := true
		for _, q := range factors {
			if modExp(g, (p-1)/q, p) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
	return 0
}

func generateSecretKey(p uint64) uint64 {
	key, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(p-2))
	return key.Uint64() + 1
}

func modExp(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

func modInverse(a, p uint64) (uint64, error) {
	g, x, _ := gcdExtended(int64(a), int64(p))
	if g != 1 {
		return 0, fmt.Errorf("no inverse exists")
	}
	return uint64((x%int64(p) + int64(p)) % int64(p)), nil
}

func gcdExtended(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := gcdExtended(b, a%b)
	return g, y1, x1 - (a/b)*y1
}
