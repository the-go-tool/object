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

func TestObject_GetIndex_Json(t *testing.T) {
	source := `[3, 2, { "a": 1, "b": 2 }]`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	t.Run("slice index get", func(t *testing.T) {
		obj := object.GetIndex(1)
		if obj.val.Float() != 2 {
			t.Fatalf(`expect 2, got: %v`, obj.val.Float())
		}
	})

	// Go isn't guaranteeing that map fields will be in same order as serialized
	t.Run("map index get unsupported", func(t *testing.T) {
		obj := object.GetIndex(2).GetIndex(1)
		if obj.GetError().Error() != ErrorTypeNotSupport {
			t.Fatalf(`maps mustn't be supported`)
		}
	})

	t.Run("out of range slice", func(t *testing.T) {
		obj := object.GetIndex(3)
		if obj.GetError().Error() != ErrorIndexRange {
			t.Fatalf(`expect out of range error`)
		}
	})
}

func TestObject_GetKeys_Json(t *testing.T) {
	source := `[3, 2, { "a": 1, "b": 2 }]`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	t.Run("array keys", func(t *testing.T) {
		keys := object.GetKeys()
		if !reflect.DeepEqual(keys, []string{"0", "1", "2"}) {
			t.Fatalf(`expect [0 1 2], got: %v`, keys)
		}
	})

	t.Run("map keys", func(t *testing.T) {
		keys := object.Get("2").GetKeys()
		if !reflect.DeepEqual(keys, []string{"a", "b"}) && !reflect.DeepEqual(keys, []string{"b", "a"}) {
			t.Fatalf(`expect [a b] or reversed, got: %v`, keys)
		}
	})
}

func TestObject_GetValues_Json(t *testing.T) {
	source := `{"a": [3, 2, 1], "b": { "c":1, "d":2 }}`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	t.Run("array values", func(t *testing.T) {
		objs := object.Get("a").GetValues()
		control := []float64{3, 2, 1}
		result := make([]float64, 0, 8)
		for _, obj := range objs {
			if obj.IsFloat() {
				result = append(result, obj.val.Float())
			}
		}
		if !reflect.DeepEqual(control, result) {
			t.Fatalf(`expect %v, got: %v`, control, result)
		}
	})

	t.Run("map values", func(t *testing.T) {
		objs := object.Get("b").GetValues()
		control := []float64{1, 2}
		result := make([]float64, 0, 8)
		for _, obj := range objs {
			if obj.IsFloat() {
				result = append(result, obj.val.Float())
			}
		}
		if !reflect.DeepEqual(control, result) && !reflect.DeepEqual(control, []float64{result[1], result[0]}) {
			t.Fatalf(`expect %v or reversed, got: %v`, control, result)
		}
	})
}

func TestObject_GetEntries_Json(t *testing.T) {
	source := `{"a": [3, 2, 1], "b": { "c":1, "d":2 }}`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	type entr struct {
		Key   string
		Value float64
	}

	t.Run("array entries", func(t *testing.T) {
		entries := object.Get("a").GetEntries()
		control := []entr{{"0", 3}, {"1", 2}, {"2", 1}}
		result := make([]entr, 0, 8)
		for _, entry := range entries {
			if entry.Value.IsFloat() {
				result = append(result, entr{
					entry.Key,
					entry.Value.val.Float(),
				})
			}
		}
		if !reflect.DeepEqual(control, result) {
			t.Fatalf(`expect %v, got: %v`, control, result)
		}
	})

	t.Run("map entries", func(t *testing.T) {
		entries := object.Get("b").GetEntries()
		control := []entr{{"c", 1}, {"d", 2}}
		result := make([]entr, 0, 8)
		for _, entry := range entries {
			if entry.Value.IsFloat() {
				result = append(result, entr{
					entry.Key,
					entry.Value.val.Float(),
				})
			}
		}
		if !reflect.DeepEqual(control, result) && !reflect.DeepEqual(control, []entr{result[1], result[0]}) {
			t.Fatalf(`expect %v or reversed, got: %v`, control, result)
		}
	})
}
