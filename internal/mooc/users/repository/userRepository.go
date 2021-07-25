package users

import (
	"errors"

	"github.com/EddCode/twitter-clone/cmd/config"
	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/EddCode/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Singup(user *models.SingupUser) (*models.User, error)
    Login(user *models.UserLogin)(*models.UserToken, error)
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

func (repo *ServiceRepository) Login(userLogin *models.UserLogin) (*models.UserToken, error) {

    if userLogin.Email == " " || len(userLogin.Password) < 6 {
        return nil, errors.New("Email/Password are incorect")
    }


    user, foundUser, _ := repo.Repository.IsUserExist(userLogin.Email)

    if foundUser == false {
        return nil, errors.New("Password / Email was wrong")
    }

    setting, errSetting := config.GetConfig()

    if errSetting != nil {
        return nil, errSetting
    }

    tokenSigned, errToken := utils.BuildJWT(user, setting.Token.Secret)

    if errToken != nil {
        return nil, errToken
    }

    token := &models.UserToken{ Token: tokenSigned}

    return token, nil
}

func (repo *ServiceRepository) Find() {

}

func (repo *ServiceRepository) Store() {

}
