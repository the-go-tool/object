package object

import (
	"testing"
)

func TestObject_IsExists(t *testing.T) {
	document := map[string]interface{}{}
	document["nil_field"] = nil
	document["field"] = "value"
	object := New(document)

	t.Run("not exists field", func(t *testing.T) {
		if object.Get("random_field").IsExists() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("nil field not exists check", func(t *testing.T) {
		if object.Get("nil_field").Get("not_exist").IsExists() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("nil field exists check", func(t *testing.T) {
		if !object.Get("nil_field").IsExists() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("index not exists at string", func(t *testing.T) {
		if object.Get("field").Index(50).IsExists() {
			t.Fatalf(`expect false`)
		}
	})
}

func TestObject_IsNil(t *testing.T) {
	document := map[string]interface{}{}
	document["nil_field"] = nil
	document["field"] = "value"
	object := New(document)

	t.Run("nil field", func(t *testing.T) {
		if !object.Get("nil_field").IsNil() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("non-nil field", func(t *testing.T) {
		if object.Get("field").IsNil() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("not exist field", func(t *testing.T) {
		if object.Get("not exist field").IsNil() {
			t.Fatalf(`expect false`)
		}
	})
}

func TestObject_IsEmpty(t *testing.T) {
	document := map[string]interface{}{}
	document["empty_number"] = 0
	document["non_empty_number"] = 5
	document["empty_string"] = ""
	document["non_empty_string"] = "value"
	document["empty_slice"] = []string{}
	document["non_empty_slice"] = []string{"value"}
	document["non_empty_slice_2"] = []string{""}
	object := New(document)

	t.Run("empty number", func(t *testing.T) {
		if !object.Get("empty_number").IsEmpty() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("non empty number", func(t *testing.T) {
		if object.Get("non_empty_number").IsEmpty() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		if !object.Get("empty_string").IsEmpty() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("non empty string", func(t *testing.T) {
		if object.Get("non_empty_string").IsEmpty() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		if !object.Get("empty_slice").IsEmpty() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("non empty slice", func(t *testing.T) {
		if object.Get("non_empty_slice").IsEmpty() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("non empty slice 2", func(t *testing.T) {
		if object.Get("non_empty_slice_2").IsEmpty() {
			t.Fatalf(`expect false`)
		}
	})
}

func TestObject_IsIntStrict(t *testing.T) {
	document := map[string]interface{}{}
	document["int1"] = int(5)
	document["non_int1"] = float32(5.1)
	document["non_int2"] = "value"
	document["non_int3"] = "5.1"
	document["non_int4"] = float32(5.0)
	document["non_int5"] = "5"
	object := New(document)

	t.Run("int number", func(t *testing.T) {
		if !object.Get("int1").IsIntStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("float number casting to int with losses", func(t *testing.T) {
		if object.Get("non_int1").IsIntStrict() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("string non-number", func(t *testing.T) {
		if object.Get("non_int2").IsIntStrict() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("string which can be casted", func(t *testing.T) {
		if object.Get("non_int3").IsIntStrict() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("float number casting to int lossless", func(t *testing.T) {
		if object.Get("non_int4").IsIntStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string which can be casted lossless", func(t *testing.T) {
		if object.Get("non_int5").IsIntStrict() {
			t.Fatalf(`expect true`)
		}
	})
}

func TestObject_IsInt(t *testing.T) {
	document := map[string]interface{}{}
	document["int1"] = int(5)
	document["int2"] = float32(5.0)
	document["int3"] = "5"
	document["non_int1"] = float32(5.1)
	document["non_int2"] = "value"
	document["non_int3"] = "5.1"
	object := New(document)

	t.Run("int number", func(t *testing.T) {
		if !object.Get("int1").IsInt() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("float number casting to int lossless", func(t *testing.T) {
		if !object.Get("int2").IsInt() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string which can be casted lossless", func(t *testing.T) {
		if !object.Get("int3").IsInt() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("float number casting to int with losses", func(t *testing.T) {
		if object.Get("non_int1").IsInt() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("string non-number", func(t *testing.T) {
		if object.Get("non_int2").IsInt() {
			t.Fatalf(`expect false`)
		}
	})

	t.Run("string which can be casted", func(t *testing.T) {
		if object.Get("non_int3").IsInt() {
			t.Fatalf(`expect false`)
		}
	})
}

func TestObject_IsFloatStrict(t *testing.T) {
	document := map[string]interface{}{}
	document["float1"] = float32(5)
	document["float2"] = float32(5.1)
	document["non_float1"] = "5"
	document["non_float2"] = "5.1"
	document["non_float3"] = int(5)
	object := New(document)

	t.Run("float number which can be float", func(t *testing.T) {
		if !object.Get("float1").IsFloatStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("float number which can't be float", func(t *testing.T) {
		if !object.Get("float2").IsFloatStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string number which can be float", func(t *testing.T) {
		if object.Get("non_float1").IsFloatStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string number which can't be float", func(t *testing.T) {
		if object.Get("non_float2").IsFloatStrict() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("int number which can be float", func(t *testing.T) {
		if object.Get("non_float3").IsFloatStrict() {
			t.Fatalf(`expect true`)
		}
	})
}

func TestObject_IsFloat(t *testing.T) {
	document := map[string]interface{}{}
	document["float1"] = float32(5)
	document["float2"] = float32(5.1)
	document["float3"] = "5"
	document["float4"] = "5.1"
	document["float5"] = int(5)
	document["non_float6"] = "value"
	object := New(document)

	t.Run("float number which can be float", func(t *testing.T) {
		if !object.Get("float1").IsFloat() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("float number which can't be float", func(t *testing.T) {
		if !object.Get("float2").IsFloat() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string number which can be float", func(t *testing.T) {
		if !object.Get("float3").IsFloat() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string number which can't be float", func(t *testing.T) {
		if !object.Get("float4").IsFloat() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("int number which can be float", func(t *testing.T) {
		if !object.Get("float5").IsFloat() {
			t.Fatalf(`expect true`)
		}
	})

	t.Run("string which can't be float", func(t *testing.T) {
		if object.Get("non_float6").IsFloat() {
			t.Fatalf(`expect false`)
		}
	})
}
