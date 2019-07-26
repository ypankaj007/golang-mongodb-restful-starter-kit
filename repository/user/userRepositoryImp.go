package user

import (
	"context"

	"go-restapis/config"
	"go-restapis/model"

	mgo "gopkg.in/mgo.v2"
)

type UserRepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func New(db *mgo.Session) UserRepository {
	return &UserRepositoryImp{db: db}
}

func (service *UserRepositoryImp) Create(ctx context.Context, user *model.User) error {
	return service.collection().Insert(user)
}

func (service *UserRepositoryImp) FindAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (service *UserRepositoryImp) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserRepositoryImp) FindOneById(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}

func (service *UserRepositoryImp) Delete(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserRepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.User, error) {
	var user model.User
	e := service.collection().Find(query).One(&user)
	return &user, e
}

func (service *UserRepositoryImp) collection() *mgo.Collection {
	return service.db.DB("go-restapis").C("user")
}
