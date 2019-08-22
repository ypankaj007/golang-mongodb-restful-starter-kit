package auth

import (
	"context"

	"golang-mongodb-restful-starter-kit/model"

	repository "golang-mongodb-restful-starter-kit/repository/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthServiceImp struct {
	db         *mgo.Session
	repository repository.UserRepository
}

func New(db *mgo.Session) AuthService {
	return &AuthServiceImp{db: db, repository: repository.New(db)}
}

func (service *AuthServiceImp) Create(ctx context.Context, user *model.User) error {

	return service.repository.Create(ctx, user)
}

func (service *AuthServiceImp) Login(ctx context.Context, credential *model.Credential) (*model.User, error) {
	query := bson.M{"email": credential.Email}
	user, err := service.repository.FindOne(ctx, query)
	if err != nil || user == nil {
		return nil, err
	}

	if err = user.ComparePassword(credential.Password); err != nil {
		return nil, err
	}
	return user, nil

}
