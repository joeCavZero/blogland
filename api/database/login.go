package database

import "fmt"

func CreateTokenByCredentials(email string, password string) (string, error) {
	if email == "joao@" && password == "123456" {
		return "valid-token-123456", nil
	}
	return "", fmt.Errorf("invalid credentials")
}
