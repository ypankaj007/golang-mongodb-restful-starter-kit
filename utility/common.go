package utility

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func CurrentTimeInMilli() int64 {
	return UnixMilli(time.Now())
}

// ****************************************************  Bson *******************************************
// Custom Bson operation

type CustomBson struct{}

type BsonWrapper struct {
	Set      interface{} `json:"$set,omitempty" bson:"$set,omitempty"`
	Unset    interface{} `json:"$unset,omitempty" bson:"$unset,omitempty"`
	Push     interface{} `json:"$push,omitempty" bson:"$push,omitempty"`
	AddToSet interface{} `json:"$addToSet,omitempty" bson:"$addToSet,omitempty"`
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var stringInterfaceMap map[string]interface{}
	itr, _ := bson.Marshal(s)
	err := bson.Unmarshal(itr, &stringInterfaceMap)
	return stringInterfaceMap, err
}

func (customBson *CustomBson) Set(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Set: data}
	return StructToMap(s)
}

func (customBson *CustomBson) Push(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Push: data}
	return StructToMap(s)
}

func (customBson *CustomBson) Unset(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{Unset: data}
	return StructToMap(s)
}

func (customBson *CustomBson) AddToSet(data interface{}) (map[string]interface{}, error) {
	s := BsonWrapper{AddToSet: data}
	return StructToMap(s)
}

// ****************************************************  Bson End  *******************************************
