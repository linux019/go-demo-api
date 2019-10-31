package util

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	r, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(r), err
}

func ComparePasswordAndHash(password, encodedHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encodedHash), []byte(password)) == nil
}
