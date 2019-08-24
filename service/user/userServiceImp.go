package user

import (
	"context"
	"log"

	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/model"

	repository "golang-mongodb-restful-starter-kit/repository/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserServiceImp struct {
	db         *mgo.Session
	repository repository.UserRepository
	config     *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) UserService {
	return &UserServiceImp{db: db, config: c, repository: repository.New(db, c)}
}

func (service *UserServiceImp) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserServiceImp) Get(ctx context.Context, id string) (*model.User, error) {
	log.Println(ctx.Value("userId"))
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	return service.repository.FindOne(ctx, query)
}
