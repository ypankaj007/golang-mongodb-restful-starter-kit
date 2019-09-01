package user

import (
	"context"
	"golang-mongodb-restful-starter-kit/model"
)

type UserService interface {
	Update(context.Context, string, *model.UserUpdate) error
	Get(context.Context, string) (*model.User, error)
}
