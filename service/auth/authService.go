package auth

import (
	"context"
	"go-restapis/model"
)

type AuthService interface {
	Create(context.Context, *model.User) error
	Login(context.Context, *model.Credential) (*model.User, error)
}
