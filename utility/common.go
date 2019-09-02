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
	Set      interface{} `json:"$set,omitempty" bson:"$set,omitempty"`
	Unset    interface{} `json:"$unset,omitempty" bson:"$unset,omitempty"`
	Push     interface{} `json:"$push,omitempty" bson:"$push,omitempty"`
	AddToSet interface{} `json:"$addToSet,omitempty" bson:"$addToSet,omitempty"`
}

func ToMap(s interface{}) (map[string]interface{}, error) {
	var stringInterfaceMap map[string]interface{}
	itr, _ := bson.Marshal(s)
	err := bson.Unmarshal(itr, &stringInterfaceMap)
	return stringInterfaceMap, err
}

func (customBson *CustomBson) Set(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Set: data}
	return ToMap(s)
}

func (customBson *CustomBson) Push(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Push: data}
	return ToMap(s)
}

func (customBson *CustomBson) Unset(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Unset: data}
	return ToMap(s)
}

func (customBson *CustomBson) AddToSet(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{AddToSet: data}
	return ToMap(s)
}

// ****************************************************  Bson End  *******************************************
