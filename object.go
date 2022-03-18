package object

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v3"
	"reflect"
)

// Object - type of anything.
type Object struct {
	val *reflect.Value
	err error
}

// New - create new object from any type.
func New(obj interface{}) Object {
	val := reflect.ValueOf(obj)
	return Object{&val, nil}
}

// NewFromValue - create new object from reflect.Value
func NewFromValue(val reflect.Value) Object {
	return Object{&val, nil}
}

// NewFromData - detect and create object from any supporting data format.
// It's supports JSON, YAML, TOML and BSON.
// XML and HTML isn't supporting, so it has extra fields which isn't straight
// data containers (like tags and attrs). Use your favourite XML/HTML
// deserializer and pass the result to New.
func NewFromData(data []byte) Object {
	var obj Object
	if obj = NewFromJson(data); obj.GetError() == nil {
		return obj
	} else if obj = NewFromYaml(data); obj.GetError() == nil {
		return obj
	} else if obj = NewFromToml(data); obj.GetError() == nil {
		return obj
	} else if obj = NewFromBson(data); obj.GetError() == nil {
		return obj
	}
	return Object{nil, newError(ErrorDataParse)}
}

// NewFromJson - create new object from json bytes
func NewFromJson(data []byte) Object {
	var document interface{}
	if err := json.Unmarshal(data, &document); err != nil {
		return Object{nil, newError(ErrorDataParse)}
	}
	return New(document)
}

// NewFromYaml - create new object from yaml bytes
func NewFromYaml(data []byte) Object {
	var document interface{}
	if err := yaml.Unmarshal(data, &document); err != nil {
		return Object{nil, newError(ErrorDataParse)}
	}
	return New(document)
}

// NewFromBson - create new object from bson bytes
func NewFromBson(data []byte) Object {
	var document interface{}
	if err := bson.Unmarshal(data, &document); err != nil {
		return Object{nil, newError(ErrorDataParse)}
	}
	return New(document)
}

// NewFromToml - create new object from toml bytes
func NewFromToml(data []byte) Object {
	var document interface{}
	if err := toml.Unmarshal(data, &document); err != nil {
		return Object{nil, newError(ErrorDataParse)}
	}
	return New(document)
}
