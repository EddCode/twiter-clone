package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/dgrijalva/jwt-go"
)

func HashPassword(password string, cost int) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

    return string(bytes), err
}

func BuildJWT(user *models.User, secret string) (string, error) {
    jwtSecret := []byte(secret)

    payload := jwt.MapClaims{
        "_id": user.ID.Hex(),
        "email": user.Email,
        "name": user.FullName,
        "exp": time.Now().Add(time.Minute * 15).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    tokenSigned, err := token.SignedString(jwtSecret)

    if err != nil {
       return tokenSigned, err
    }

    return tokenSigned, nil

}
