package encript

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password :: %v", err)
	}

	return string(hash), nil

}

func CompirePassword(hasedpassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedpassword), []byte(password))
}
