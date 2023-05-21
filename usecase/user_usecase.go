package usecase

import (
	"echo-rest-api/model"
	"echo-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	SignUp(user *model.User) (*model.UserResponse, error)
	Login(user *model.User) (string, error)
}

type UserUseCaseImpl struct {
	userRepo repository.UserRepository
}

// TODO: 理解する
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{r}
}

func (uc *UserUseCaseImpl) SignUp(user *model.User) (*model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uc.userRepo.CreateUser(newUser); err != nil {
		return nil, err
	}
	resUser := &model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uc *UserUseCaseImpl) Login(user *model.User) (string, error) {
	storedUser := model.User{}
	if err := uc.userRepo.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	signedString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
