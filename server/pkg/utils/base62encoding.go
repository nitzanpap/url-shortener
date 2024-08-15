package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

const base62Characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// base62Encode encodes a byte slice into a Base62 string.
func Base62Encode(bytes []byte) string {
	var number big.Int
	number.SetBytes(bytes) // Set the big int from the byte slice

	base := big.NewInt(62)
	remainder := new(big.Int)
	var result string

	for number.Cmp(big.NewInt(0)) == 1 { // While number > 0
		number.DivMod(&number, base, remainder)
		result = string(base62Characters[remainder.Int64()]) + result
	}

	return result
}

// base62Decode decodes a Base62 string back into a byte slice.
func Base62Decode(b62 string) ([]byte, error) {
	var number big.Int
	base := big.NewInt(62)

	// Decode the Base62 string to a big.Int
	for i := 0; i < len(b62); i++ {
		index := big.NewInt(int64(indexInBase62Characters(b62[i])))
		if index.Cmp(big.NewInt(-1)) == 0 {
			return nil, fmt.Errorf("invalid character: %s", string(b62[i]))
		}
		number.Mul(&number, base)
		number.Add(&number, index)
	}

	return number.Bytes(), nil
}

// indexInBase62Characters returns the index of a byte in the base62Characters.
func indexInBase62Characters(char byte) int {
	for i, b := range []byte(base62Characters) {
		if b == char {
			return i
		}
	}
	return -1
}

func GenerateTruncatedHashInBase62(str string, numOfBytesToTruncate int) string {
	hash := sha256.Sum256([]byte(str))
	base62String := Base62Encode(hash[:])
	truncatedBase62Hash := base62String[:numOfBytesToTruncate]
	return truncatedBase62Hash
}
