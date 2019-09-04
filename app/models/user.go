package models

import (
	"golang-mongodb-restful-starter-kit/utility"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// User , definds user model
type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty" bson:"name,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty"`
	Password  string        `json:"-" bson:"password,omitempty"`
	Salt      string        `json:"-" bson:"salt,omitempty"`
	Role      string        `json:"role,omitempty" bson:"role,omitempty"`
	IsActive  bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
	CreatedAT int64         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAT int64         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// Credential , definds login credential model
type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserUpdate , definds user update model
type UserUpdate struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive bool   `json:"isActive,omitempty" bson:"isActive,omitempty"`
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

// ComparePassword , used to compared
// hashed password with input text password
// return error if any otherwise nil
func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

// Initialize , will set the hashed password, createdAt and updatedAt
// date in milliseconds
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
	u.IsActive = true
	u.Role = utility.UserRole
	return nil
}

// Validate user fields
// This function validates user data
// and return error is any
// all errors are related to the fields
func (u *User) Validate() error {

	// validating name field with retuired, min length 3, max length 25 and no regex check
	if e := utility.ValidateRequireAndLengthAndRegex(u.Name, true, 3, 25, "", "Name"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := utility.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "Email"); e != nil {
		return e
	}

	// validating password field with required, min length 8, max length 25 and regex check
	if e := utility.ValidateRequireAndLengthAndRegex(u.Password, true, 8, 25, "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$", "Password"); e != nil {
		return e
	}

	return nil

}
