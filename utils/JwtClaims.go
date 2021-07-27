package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
    Email string `json:"email"`
    ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
    jwt.StandardClaims
}
