package auth

import (
	"context"

	model "golang-mongodb-restful-starter-kit/app/models"

	repository "golang-mongodb-restful-starter-kit/app/repositories/user"

	"gopkg.in/mgo.v2/bson"
)

type AuthServiceInterface interface {
	Create(context.Context, *model.User) error
	Login(context.Context, *model.Credential) (*model.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}

type AuthService struct {
	repository repository.UserRepository
}

func New(userRepo repository.UserRepository) AuthServiceInterface {
	return &AuthService{repository: userRepo}
}

func (service *AuthService) Create(ctx context.Context, user *model.User) error {

	return service.repository.Create(ctx, user)
}

func (service *AuthService) Login(ctx context.Context, credential *model.Credential) (*model.User, error) {
	query := bson.M{"email": credential.Email}
	user, err := service.repository.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(credential.Password); err != nil {
		return nil, err
	}
	return user, nil

}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *AuthService) IsUserAlreadyExists(ctx context.Context, email string) bool {

	return service.repository.IsUserAlreadyExists(ctx, email)

}
