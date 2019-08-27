package user

import (
	"context"
	"golang-mongodb-restful-starter-kit/model"
)

type UserService interface {
	Update(context.Context, *model.User) error
	Get(context.Context, string) (*model.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}
