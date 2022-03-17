package object

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func ExampleObject_Get() {
	//source := `{
	//	"a": "value",
	//	"b": {
	//		"c": -500.5
	//	},
	//	"d": [3, 2, 1]
	//}`
	//var document interface{}
	//_ = json.Unmarshal([]byte(source), &document)
	//object := New(document)
	//
	//object.Get("a").String() // "value"
}

func TestObject_Get_Json(t *testing.T) {
	source := `{
		"a": "value",
		"b": {
			"c": -500.5
		},
		"d": [3, 2, 1]
	}`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	t.Run("non-nil check", func(t *testing.T) {
		obj := object.Get("a")
		if obj.val == nil {
			t.Fatalf(`unexpected nil`)
		}
	})

	t.Run("field not found check", func(t *testing.T) {
		obj := object.Get("not-exists").Get("fields").Get("chain")
		if errors.Is(obj.GetError(), errors.New(ErrorFieldNotFound)) {
			t.Fatalf(`expect ErrorFieldNotFound error, got: %v`, obj.GetError())
		}
	})

	t.Run("string kind check", func(t *testing.T) {
		obj := object.Get("a")
		if obj.val.Kind() != reflect.String {
			t.Fatalf(`expect kind string, got: %v`, obj.val.Kind())
		}
	})

	t.Run("get value", func(t *testing.T) {
		obj := object.Get("a")
		if obj.val.String() != "value" {
			t.Fatalf(`expect "value", got: %v`, obj.val.String())
		}
	})

	t.Run("get value sequential", func(t *testing.T) {
		obj := object.Get("b").Get("c")
		if obj.val.Float() != -500.5 {
			t.Fatalf(`expect -500.5, got: %v`, obj.val.String())
		}
	})

	t.Run("get slice element", func(t *testing.T) {
		obj := object.Get("d").Get("1")
		if obj.val.Float() != 2 {
			t.Fatalf(`expect 2, got: %v`, obj.val.Float())
		}
	})
}
