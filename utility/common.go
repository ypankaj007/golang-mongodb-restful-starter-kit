package utility

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SetWraper struct {
	Set interface{} `json:"$set,omitempty" bson:"$set,omitempty"`
}

func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func CurrentTimeInMilli() int64 {
	return UnixMilli(time.Now())
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var stringInterfaceMap map[string]interface{}
	itr, _ := bson.Marshal(s)
	err := bson.Unmarshal(itr, &stringInterfaceMap)
	return stringInterfaceMap, err
}

func SetDataToBson(data interface{}) (map[string]interface{}, error) {
	s := SetWraper{Set: data}
	return StructToMap(s)
}
