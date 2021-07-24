package users

import (
	"errors"

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
	Repository *Repository
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &ServiceRepository{ Repository: NewRepository(db) }
}

func (repo *ServiceRepository) Singup(user *models.SingupUser) (*models.User, error) {

    _, errExist, _ := repo.Repository.IsUserExist(user.Email)

    if errExist {
        return nil, errors.New("User already exist")
    }

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

    _, err := repo.Repository.StoreUser(newUser)

    if err != nil {
        return nil, err
    }

    return newUser, nil
}

func (repo *ServiceRepository) Find() {

}

func (repo *ServiceRepository) Store() {

}
