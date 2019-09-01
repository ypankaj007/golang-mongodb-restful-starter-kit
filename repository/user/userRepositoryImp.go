package user

import (
	"context"

	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) UserRepository {
	return &UserRepositoryImp{db: db, config: c}
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
	var user model.User
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	e := service.collection().Find(query).Select(bson.M{"password": 0, "salt": 0}).One(&user)
	return &user, e
}

func (service *UserRepositoryImp) Delete(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserRepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.User, error) {
	var user model.User
	e := service.collection().Find(query).One(&user)
	return &user, e
}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *UserRepositoryImp) IsUserAlreadyExists(ctx context.Context, email string) bool {
	query := bson.M{"email": email}
	_, e := service.FindOne(ctx, query)
	if e != nil {
		return false
	}
	return true
}

func (service *UserRepositoryImp) collection() *mgo.Collection {
	return service.db.DB(service.config.DataBaseName).C("users")
}
