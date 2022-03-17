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

// IsNumberStrict - check that the object value is one of number type.
// It includes all integer and float types. To determine which
// one - use IsIntStrict or IsFloatStrict
func (o Object) IsNumberStrict() bool {
	return o.IsIntStrict() || o.IsFloatStrict()
}

// IsNumber - check that the object value is one of number type or can be cast.
// It includes all integer and float types. To determine which
// one - use IsInt or IsFloat
func (o Object) IsNumber() bool {
	return o.IsInt() || o.IsFloat()
}
