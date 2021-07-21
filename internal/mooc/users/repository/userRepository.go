package users

import (
	"errors"
	"log"

	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Singup(user *models.SingupUser) (*models.User, error)
	Store()
	Find()
}

type ServiceRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &ServiceRepository{db: db}
}

func (repo *ServiceRepository) Singup(user *models.SingupUser) (*models.User, error) {
	log.Printf("execute singup repository function user info %v", user)

    if len(user.Email) == 0 {
        return nil, errors.New("Missing email")
    }

    if len(user.Password) < 6 {
        return nil, errors.New("Password has to be more than y characters")
    }

    newUser := &models.User{
        ID: primitive.NewObjectID(),
        Avatar: "no avatar",
        Biography: "Biography",
        Location: "Location",
        FullName: user.FullName,
        Birthday: user.Birthday,
        Phone: user.Phone,
        Email: user.Email,
        Password: user.Password,
    }

    return newUser, nil
}

func (repo *ServiceRepository) Find() {

}

func (repo *ServiceRepository) Store() {

}
