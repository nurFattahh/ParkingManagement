package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashValue(value string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	return string(bytes), err

}

func CompareHash(value string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil

}
