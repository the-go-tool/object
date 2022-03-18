package object

import (
	"reflect"
	"strconv"
)

// IsExists - check that the object exists.
func (o Object) IsExists() bool {
	return o.val != nil
}

// IsNil - check that the object value is nil/null.
func (o Object) IsNil() bool {
	if !o.IsExists() {
		return false
	}
	switch o.val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Invalid:
		if !o.val.IsValid() || o.val.IsNil() {
			return true
		}
	}
	return false
}

// IsEmpty - check that the object value empty.
// reflect.Array isn't contain any elements and reflect.Struct isn't contain any fields.
// Other types is equal their default values.
func (o Object) IsEmpty() bool {
	if !o.IsExists() {
		return false
	}
	switch o.val.Kind() {
	case reflect.Array, reflect.Slice:
		return o.val.Len() == 0
	case reflect.Struct:
		return o.val.NumField() == 0
	}
	return o.IsExists() && o.val.IsZero()
}

// IsMap - check that the object value is map.
func (o Object) IsMap() bool {
	return o.IsExists() && o.val.Kind() == reflect.Map
}

// IsSlice - check that the object value is slice.
func (o Object) IsSlice() bool {
	return o.IsExists() && o.val.Kind() == reflect.Slice
}

// IsIntStrict - check that the object is integer number.
func (o Object) IsIntStrict() bool {
	if !o.IsExists() {
		return false
	}
	switch o.val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

// IsInt - check that the object is integer or can be cast lossless.
func (o Object) IsInt() bool {
	if !o.IsExists() {
		return false
	}
	if o.IsIntStrict() {
		return true
	}
	switch o.val.Kind() {
	case reflect.Float32, reflect.Float64:
		if float64(int64(o.val.Float())) == o.val.Float() {
			return true
		}
	case reflect.String:
		num, err := strconv.ParseFloat(o.val.String(), 64)
		if err != nil {
			return false
		}
		if float64(int64(num)) == num {
			return true
		}
	}
	return false
}

// IsFloatStrict - check that the object is float number.
func (o Object) IsFloatStrict() bool {
	if !o.IsExists() {
		return false
	}
	switch o.val.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// IsFloat - check that the object is float or can be cast.
func (o Object) IsFloat() bool {
	if !o.IsExists() {
		return false
	}
	if o.IsFloatStrict() || o.IsInt() {
		return true
	}
	switch o.val.Kind() {
	case reflect.String:
		_, err := strconv.ParseFloat(o.val.String(), 64)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

// IsStringStrict - check that the object is string.
func (o Object) IsStringStrict() bool {
	return o.IsExists() && o.val.Kind() == reflect.Map
}

// IsString - check that the object is string or can be cast.
// Simple types like numbers always can be cast. If you would like
// to get something more complicated - use ToJson or similar serialization.
// If the required type serialization isn't presented - use ToValue and
// pass it to your favourite serializer manually.
func (o Object) IsString() bool {
	if !o.IsExists() {
		return false
	}
	if o.IsStringStrict() || o.IsFloatStrict() {
		return true
	}
	return false
}

// IsBoolStrict - check that the object is boolean.
func (o Object) IsBoolStrict() bool {
	return o.IsExists() && o.val.Kind() == reflect.Bool
}

// IsBool - check that the object is boolean or can be cast.
// For string truthy values are: "true", "yes", "on"
// and if it can be cast to number - any except zero.
// For numbers truthy values any except zero.
//TODO: Make the realisation and tests
func (o Object) IsBool() bool {
	if !o.IsExists() {
		return false
	}
	if o.IsBoolStrict() {
		return true
	}
	switch o.val.Kind() {
	case reflect.Float32, reflect.Float64:
		if float64(int64(o.val.Float())) == o.val.Float() {
			return true
		}
	case reflect.String:
		num, err := strconv.ParseFloat(o.val.String(), 64)
		if err != nil {
			return false
		}
		if float64(int64(num)) == num {
			return true
		}
	}
	return false
}
