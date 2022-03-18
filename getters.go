package object

import (
	"fmt"
	"reflect"
	"strconv"
)

// Entry - object that represents a key-value pair for such methods
// like GetEntries, ForEach, Map, etc.
type Entry struct {
	Key   string
	Value Object
}

// GetError - last operation error info.
func (o Object) GetError() error {
	return o.err
}

// Get - get sub-object by their key-name.
// Can be applied to reflect.Map and reflect.Slice kinds.
// Acts like GetIndex if Object is reflect.Slice.
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
		return o.GetIndex(int(index))
	}

	return Object{nil, newError(ErrorTypeNotSupport)}
}

// GetIndex - get sub-object by their index in slice.
func (o Object) GetIndex(index int) Object {
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

// GetKeys - get keys names of reflect.Map or indexes for reflect.Slice.
// For reflect.Map order of keys is not guaranteeing.
func (o Object) GetKeys() []string {
	if !o.IsExists() {
		return []string{}
	}
	keys := make([]string, 0, 16)

	val := deref(*o.val)
	switch val.Kind() {
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			keys = append(keys, iter.Key().String())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			keys = append(keys, fmt.Sprintf("%d", i))
		}
	}

	return keys
}

// GetValues - get values of reflect.Map or reflect.Slice.
// For reflect.Map order of values is not guaranteeing.
func (o Object) GetValues() []Object {
	if !o.IsExists() {
		return []Object{}
	}
	values := make([]Object, 0, 16)

	val := deref(*o.val)
	switch val.Kind() {
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			val := iter.Value().Elem()
			values = append(values, Object{&val, nil})
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			val := val.Index(i).Elem()
			values = append(values, Object{&val, nil})
		}
	}

	return values
}

// GetEntries - get key-values of reflect.Map or reflect.Slice.
// For reflect.Map order of entries is not guaranteeing.
func (o Object) GetEntries() []Entry {
	if !o.IsExists() {
		return []Entry{}
	}
	entries := make([]Entry, 0, 16)

	val := deref(*o.val)
	switch val.Kind() {
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			val := iter.Value().Elem()
			entries = append(entries, Entry{
				Key:   iter.Key().String(),
				Value: Object{&val, nil},
			})
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			val := val.Index(i).Elem()
			entries = append(entries, Entry{
				Key:   fmt.Sprintf("%d", i),
				Value: Object{&val, nil},
			})
		}
	}

	return entries
}
