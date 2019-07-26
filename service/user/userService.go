package user

import (
	"context"
	"go-restapis/model"
)

type UserService interface {
	Update(context.Context, *model.User) error
	Get(context.Context, string) (*model.User, error)
}
