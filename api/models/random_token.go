package models

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomToken() (string, error) {
	const tokenLength = 128
	randomBytes := make([]byte, tokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(randomBytes)

	return token, nil
}
