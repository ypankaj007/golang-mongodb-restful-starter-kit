package auth

import (
	"context"
	"golang-mongodb-restful-starter-kit/model"
)

type AuthService interface {
	Create(context.Context, *model.User) error
	Login(context.Context, *model.Credential) (*model.User, error)
}
