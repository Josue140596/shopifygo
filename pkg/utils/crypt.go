package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CryptPassword(password string) (string, error) {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash : " + err.Error())
	}

	return string(hashedPassword), nil
}
