package db

import (
	"golang-mongodb-restful-starter-kit/config"

	mgo "gopkg.in/mgo.v2"
)

var instanace *mgo.Session

var err error

// GetInstance return copy of db session
func GetInstance(c *config.Configuration) *mgo.Session {

	if instanace == nil {
		instanace, err = mgo.Dial(c.DataBaseConnectionURL)
		if err != nil {
			panic(err)
		}
	}
	return instanace.Copy()

}
