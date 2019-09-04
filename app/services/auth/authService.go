package auth

import (
	"context"
	model "golang-mongodb-restful-starter-kit/app/models"
)

type AuthService interface {
	Create(context.Context, *model.User) error
	Login(context.Context, *model.Credential) (*model.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}
