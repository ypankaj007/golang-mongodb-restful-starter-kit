package model

import (
	"golang-mongodb-restful-starter-kit/utility"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	Salt      string        `json:"salt" bson:"salt"`
	Role      string        `json:"role" bson:"role"`
	IsActive  bool          `json:"isActive" bson:"isActive"`
	CreatedAT int64         `json:"createdAt" bson:"createdAt"`
	UpdatedAT int64         `json:"updatedAt" bson:"updatedAt"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func UserIndex() mgo.Index {
// 	return mgo.Index{
// 		Key:        []string{"email"},
// 		Unique:     true,
// 		DropDups:   true,
// 		Background: true,
// 		Sparse:     true,
// 	}
// }

func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordBytes := []byte(u.Password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash[:])
	u.Salt = salt
	u.CreatedAT = utility.CurrentTimeInMilli()
	u.UpdatedAT = utility.CurrentTimeInMilli()
	u.Role = utility.UserRole
	return nil
}
