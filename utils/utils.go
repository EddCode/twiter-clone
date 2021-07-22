package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, cost int) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

    return string(bytes), err
}
