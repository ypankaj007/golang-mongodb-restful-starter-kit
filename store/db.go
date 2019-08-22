package store

import (
	"golang-mongodb-restful-starter-kit/config"
	"sync"
	"sync/atomic"

	mgo "gopkg.in/mgo.v2"
)

var mux sync.Mutex

var instanace *mgo.Session

var initialized uint32
var err error

// GetInstance return copy of db session
func GetInstance(c *config.Configuration) *mgo.Session {

	if atomic.LoadUint32(&initialized) == 1 {
		return instanace.Copy()
	}

	mux.Lock()
	defer mux.Unlock()

	if instanace == nil {
		instanace, err = mgo.Dial(c.DatabaseConnectionURL)
		if err != nil {
			panic(err)
		}
		atomic.StoreUint32(&initialized, 1)
	}
	return instanace.Copy()

}
