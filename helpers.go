package object

import "reflect"

// deref - acts like *operator but in deep mode.
func deref(v reflect.Value) reflect.Value {
	cpv := v
	for cpv.Kind() == reflect.Ptr {
		cpv = cpv.Elem()
	}
	return cpv
}
