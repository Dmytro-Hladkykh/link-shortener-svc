package data

import (
	"crypto/rand"
	"math/big"
)

const (
    alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    length   = 6
)

func GenerateShortCode() (string, error) {
    shortCode := make([]byte, length)
    alphabetLen := big.NewInt(int64(len(alphabet)))

    for i := 0; i < length; i++ {
        index, err := rand.Int(rand.Reader, alphabetLen)
        if err != nil {
            return "", err
        }
        shortCode[i] = alphabet[index.Int64()]
    }

    return string(shortCode), nil
}