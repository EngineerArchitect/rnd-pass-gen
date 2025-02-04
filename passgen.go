package main

import (
	"crypto/rand"
	"math/big"
)

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	allChars     = lowercase + uppercase + digits + specialChars
)

func generatePassword(length int, useSpecial bool) (string, error) {
	var charset string
	if useSpecial {
		charset = allChars
	} else {
		charset = lowercase + uppercase + digits
	}

	password := make([]byte, length)
	for i := range password {
		randInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randInt.Int64()]
	}

	return string(password), nil
}
