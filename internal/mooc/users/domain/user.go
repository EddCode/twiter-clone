package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName  string             `bson:"name" json:"name,omitempty"`
	Phone     string             `bson:"phone" json:"phone"`
	Birthday  time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Banner    string             `bson:"banner" json:"banner"`
	Biography string             `bson:"biography" json:"biography"`
	Location  string             `bson:"location" json:"location"`
}

type SingupUser struct {
	FullName string    `bson:"name" json:"name,omitempty"`
	Phone    string    `bson:"phone" json:"phone"`
	Birthday time.Time `bson:"birthday" json:"birthday,omitempty"`
	Email    string    `bson:"email" json:"email"`
	Password string    `bson:"password,omitempty" json:"password,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserToken struct {
	Token string `json:"token,omitempty"`
}
