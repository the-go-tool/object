package object

import (
	"reflect"
	"strconv"
)

// Get - get sub-object by their key-name.
// Can be applied to reflect.Map and reflect.Slice kinds.
// Acts like Index if Object is reflect.Slice.
func (o Object) Get(key string) Object {
	if !o.IsExists() {
		return Object{nil, newError(ErrorObjectNotExists)}
	}

	val := deref(*o.val)
	switch val.Kind() {
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			if iter.Key().String() == key {
				v := iter.Value().Elem()
				return Object{&v, nil}
			}
		}
		return Object{nil, newError(ErrorFieldNotFound)}
	case reflect.Slice, reflect.Array:
		index, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return Object{nil, newError(ErrorIndexParse)}
		}
		return o.Index(int(index))
	}

	return Object{nil, newError(ErrorTypeNotSupport)}
}
