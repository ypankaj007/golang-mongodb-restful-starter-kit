/** utility
Author - Pankaj Kr Yadav
Files contains operations related time and Bson
Sometimes we need to format or convert time, all functions related to this will be write
in this file.
These are also Bson related operations,
like convert struct to Bson etc will be in this file
*/
package utility

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UnixMilli use to get milliseconds of given time
// @params - time
// return - milliseconds of the given time
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// CurrentTimeInMilli use to get current time in milliseconds
// This function will use when we need current timestamp
// This functions return current timestamp (in millisecods)
func CurrentTimeInMilli() int64 {
	return UnixMilli(time.Now())
}

// ****************************************************  Bson *******************************************
// Custom Bson operation

// CustomBson use to perform custom bson related operatios
// like set, push, unset etc by using structs
// This will very usefull when we need to create bson map from struct
type CustomBson struct{}

// BsonWrapper contains basic bson operations
// like $set, $push, $addToSet
// It is very usefull to convert struct in bson
type BsonWrapper struct {

	// Set will to set data in db
	// example - if it required to set "name":"Jack", then it
	// need to create an struct whick contains name field and assign that
	// struct in this field. After encoded in bson it will be like
	// { $set : {name : "Jack"}} and this will usefull in mongo query
	Set interface{} `json:"$set,omitempty" bson:"$set,omitempty"`

	// The Unset operator deletes a particular field.
	// If the field does not exist, then Unset does nothing
	// If needs to unset name field then simply we will create a struct that
	// contains name field and ten set "" to name.
	// Now to unset, set that struct to Unset field. After encode that will
	// becaome like { $unset: { name: "" } }
	Unset interface{} `json:"$unset,omitempty" bson:"$unset,omitempty"`

	// The Push operator appends a specified value to an array.
	// If the field is absent in the document to update,
	// Push adds the array field with the value as its element.
	// If the field is not an array, the operation will fail.
	Push interface{} `json:"$push,omitempty" bson:"$push,omitempty"`

	// The AddToSet operator adds a value to an array unless
	// the value is already present, in which case AddToSet does nothing to that array.
	// If you use AddToSet on a field is absent in the document to update,
	// AddToSet creates the array field with the specified value as its element.
	AddToSet interface{} `json:"$addToSet,omitempty" bson:"$addToSet,omitempty"`
}

// ToMap converts interface to to map.
// It takes interface as params and
// returns map and error if any
func ToMap(s interface{}) (map[string]interface{}, error) {
	var stringInterfaceMap map[string]interface{}
	itr, _ := bson.Marshal(s)
	err := bson.Unmarshal(itr, &stringInterfaceMap)
	return stringInterfaceMap, err
}

// Set creates query to replaces the value of a field with the specified value
// params - data that need to be set
// returns - query map and error if any
func (customBson *CustomBson) Set(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Set: data}
	return ToMap(s)
}

// Push creates query to append a specified value to an array field
// params - data that need to append
// returns query map and error if any
func (customBson *CustomBson) Push(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Push: data}
	return ToMap(s)
}

// Unset creates query to delete a particular field
// params - data that need to be unset
// returns - query map and error if any
func (customBson *CustomBson) Unset(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Unset: data}
	return ToMap(s)
}

// AddToSet creates query to add a value to an array unless the value is already present.
// params - data that need to add to set,
// returns - query map and error if nay
func (customBson *CustomBson) AddToSet(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{AddToSet: data}
	return ToMap(s)
}

// ****************************************************  Bson End  *******************************************
