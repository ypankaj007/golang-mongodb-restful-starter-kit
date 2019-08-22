package user

import (
	"context"
	"golang-mongodb-restful-starter-kit/model"
)

// UserRepository , used to perform DB oprations
// Inteface contains basic oprations on user document
// So that, db opration can be perform easily
type UserRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.User) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]*model.User, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, string) (*model.User, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, *model.User) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, *model.User) error

	// FindOne, will find one entry of user matched by the query
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*model.User, error)
}
