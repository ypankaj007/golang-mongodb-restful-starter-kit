package user

import (
	"context"
	"go-restapis/model"
)

// UserRepository , used to perform DB oprations
// Inteface contains basic oprations on user document
// So that, db opration can be perform easily
type UserRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.User) error

	FindAll(context.Context) ([]*model.User, error)
	FindOneById(context.Context, string) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, *model.User) error
	FindOne(context.Context, interface{}) (*model.User, error)
}
