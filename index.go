package object

import (
	"reflect"
)

// Index - get sub-object by their index in slice.
func (o Object) Index(index int) Object {
	if !o.IsExists() {
		return Object{nil, newError(ErrorObjectNotExists)}
	}

	val := deref(*o.val)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		if index < 0 || index >= val.Len() {
			return Object{nil, newError(ErrorIndexRange)}
		}
		v := val.Index(index).Elem()
		return Object{&v, nil}
	}

	return Object{nil, newError(ErrorTypeNotSupport)}
}
