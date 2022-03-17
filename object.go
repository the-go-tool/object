package object

import (
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

// Error - last operation error info.
func (o Object) GetError() error {
	return o.err
}
