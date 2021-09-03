package users

import (
	"errors"
	"reflect"
	"strings"

	"github.com/EddCode/twitter-clone/cmd/config"
	"github.com/EddCode/twitter-clone/internal/application/customError"
	models "github.com/EddCode/twitter-clone/internal/mooc/users/domain"
	"github.com/EddCode/twitter-clone/internal/storage"
	"github.com/EddCode/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Singup(user *models.SingupUser) (*models.User, *customError.CustomError)
	Login(user *models.UserLogin) (*models.UserToken, *customError.CustomError)
	GetUserProfile(id string) (*models.User, *customError.CustomError)
	UpdateUserProfile(user models.User) (bool, *customError.CustomError)
}

type ServiceRepository struct {
	Repository *storage.Repository
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return &ServiceRepository{Repository: storage.NewRepository(db)}
}

func (repo *ServiceRepository) Singup(user *models.SingupUser) (*models.User, *customError.CustomError) {

	_, errExist, _ := repo.IsUserExist(user.Email)

	if errExist {
		return nil, customError.ThrowError("BadRequest", errors.New("User already exist"))
	}

	if len(user.Email) == 0 {
		return nil, customError.ThrowError("BadRequest", errors.New("Missing email"))
	}

	if len(user.Password) < 6 {
		return nil, customError.ThrowError("BadRequest", errors.New("Password has to be more than y characters"))
	}

	newUser := &models.User{
		ID:        primitive.NewObjectID(),
		Avatar:    "no avatar",
		Biography: "Biography",
		Location:  "Location",
		FullName:  user.FullName,
		Birthday:  user.Birthday,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
	}

	_, err := repo.StoreUser(newUser)

	if err != nil {
		return nil, customError.ThrowError("BadServerError", err)
	}

	return newUser, nil
}

func (repo *ServiceRepository) Login(userLogin *models.UserLogin) (*models.UserToken, *customError.CustomError) {

	if userLogin.Email == " " || len(userLogin.Password) < 6 {
		return nil, customError.ThrowError("Unauthorized", errors.New("Email/Password are incorect"))
	}

	user, foundUser, _ := repo.IsUserExist(userLogin.Email)

	userPwd := []byte(userLogin.Password)
	hashedPwd := []byte(user.Password)

	pwdErr := bcrypt.CompareHashAndPassword(hashedPwd, userPwd)

	if foundUser == false || pwdErr != nil {
		return nil, customError.ThrowError("Unauthorized", errors.New("Password / Email was wrong"))
	}

	setting, errSetting := config.GetConfig()

	if errSetting != nil {
		return nil, customError.ThrowError("BadServerError", errSetting)
	}

	tokenSigned, errToken := utils.BuildJWT(user, setting.Token.Secret)

	if errToken != nil {
		return nil, customError.ThrowError("Unauthorized", errToken)
	}

	token := &models.UserToken{Token: tokenSigned}

	return token, nil
}

func (repo *ServiceRepository) GetUserProfile(id string) (*models.User, *customError.CustomError) {
	profile, err := repo.getUserById(id)

	if err != nil {
		return nil, customError.ThrowError("NotFound", err)
	}

	return profile, nil

}

func (repo *ServiceRepository) UpdateUserProfile(user models.User) (bool, *customError.CustomError) {
	userProfile := make(map[string]interface{})

	fields := reflect.TypeOf(user)
	values := reflect.ValueOf(user)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		if !value.IsZero() {
			key := strings.ToLower(field.Name)
			userProfile[key] = value
		}
	}

	updated, err := repo.updateUserProfile(userProfile, utils.Claim.ID)

	if err != nil {
		return updated, customError.ThrowError("BadServerError", err)
	}

	return updated, nil
}
