package object

import (
	"encoding/json"
	"testing"
)

func TestObject_Index_Json(t *testing.T) {
	source := `[3, 2, { "a": 1, "b": 2 }]`
	var document interface{}
	_ = json.Unmarshal([]byte(source), &document)
	object := New(document)

	t.Run("slice index get", func(t *testing.T) {
		obj := object.Index(1)
		if obj.val.Float() != 2 {
			t.Fatalf(`expect 2, got: %v`, obj.val.Float())
		}
	})

	// Go isn't guaranteeing that map fields will be in same order as serialized
	t.Run("map index get unsupported", func(t *testing.T) {
		obj := object.Index(2).Index(1)
		if obj.GetError().Error() != ErrorTypeNotSupport {
			t.Fatalf(`maps mustn't be supported`)
		}
	})

	t.Run("out of range slice", func(t *testing.T) {
		obj := object.Index(3)
		if obj.GetError().Error() != ErrorIndexRange {
			t.Fatalf(`expect out of range error`)
		}
	})
}
