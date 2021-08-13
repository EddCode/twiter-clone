package utils

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/EddCode/twitter-clone/cmd/config"
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
		"_id":   user.ID.Hex(),
		"email": user.Email,
		"name":  user.FullName,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenSigned, err := token.SignedString(jwtSecret)

	if err != nil {
		return tokenSigned, err
	}

	return tokenSigned, nil
}

func ValidToken(token string) (*Claims, error) {
	setting, errSetting := config.GetConfig()

	if errSetting != nil {
		return nil, errSetting
	}

	tokenKey := []byte(setting.Token.Secret)
	claim := &Claims{}

	splitToken := strings.Split(token, "Bearera")

	if len(splitToken) != 2 {
		return nil, errors.New("invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	_, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil {
		return nil, err
	}

	return claim, nil
}
