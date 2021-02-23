package user

import (
	"context"
	"golang-mongodb-restful-starter-kit/utility"

	model "golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/config"

	repository "golang-mongodb-restful-starter-kit/app/repositories/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserServiceInterface interface {
	Update(context.Context, string, *model.UserUpdate) error
	Get(context.Context, string) (*model.User, error)
}

// UserService , implements UserService
// and perform user related business logics
type UserService struct {
	db         *mgo.Session
	repository repository.UserRepository
	config     *config.Configuration
}

// New function will initialize UserService
func New(userRepo repository.UserRepository) UserServiceInterface {
	return &UserService{repository: userRepo}
}

// Update function will update the user info
// return error if any
func (service *UserService) Update(ctx context.Context, id string, user *model.UserUpdate) error {
	query := bson.M{"_id": bson.ObjectIdHex(id), "isActive": true}
	CustomBson := &utility.CustomBson{}
	change, err := CustomBson.Set(user)
	if err != nil {
		return err
	}
	return service.repository.Update(ctx, query, change)
}

// Get function will find user by id
// return user and error if any
func (service *UserService) Get(ctx context.Context, id string) (*model.User, error) {
	return service.repository.FindOneById(ctx, id)
}
