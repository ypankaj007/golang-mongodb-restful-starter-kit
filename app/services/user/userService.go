package user

import (
	"context"
	model "golang-mongodb-restful-starter-kit/app/models"
)

type UserService interface {
	Update(context.Context, string, *model.UserUpdate) error
	Get(context.Context, string) (*model.User, error)
}
