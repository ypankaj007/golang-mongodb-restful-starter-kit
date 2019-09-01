package user

import (
	"context"

	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/model"

	repository "golang-mongodb-restful-starter-kit/repository/user"

	mgo "gopkg.in/mgo.v2"
)

// UserServiceImp , implements UserService
// and perform user related business logics
type UserServiceImp struct {
	db         *mgo.Session
	repository repository.UserRepository
	config     *config.Configuration
}

// New function will initialize UserService
// taking db session and config in params
// db session is required to perform db operations
// config is required to get the info
func New(db *mgo.Session, c *config.Configuration) UserService {
	return &UserServiceImp{db: db, config: c, repository: repository.New(db, c)}
}

// Update function will update the user info
// return error if any
func (service *UserServiceImp) Update(ctx context.Context, user *model.User) error {
	return nil
}

// Get function will find user by id
// return user and error if any
func (service *UserServiceImp) Get(ctx context.Context, id string) (*model.User, error) {
	return service.repository.FindOneById(ctx, id)
}
