# Object
> Tool to work with unspecified-scheme objects (like unmarshalled json, yaml, etc.)

> ❤️ I Need in Your Feedback! Together we'll make a better tool!  
> Please, make an **issue or PR if** you found
> **inconsistent** behavior, **errors** or **missed test-cases**.
> Or just [leave a comment at gitter](https://gitter.im/the-go-tool-object/community)
> for new feature request.

Sometimes, we can't guess a scheme for some serialized object.
By the way, we can't describe the scheme in Go structures.
But we have to work with the unknown object.
And, here is several ways to solve this trouble:

- Make a partial-filled structures for known part of the object.
Of course, it isn't common way for all cases. Especially, for cases
that require some dynamics in behavior or when some field can have
different type for different cases.
- Unmarshal the object into `interface{}` and work with them via
type-casting or use reflection. This is a hard way with tons of code.
Usually, nobody understands how it works and what happens. Even the
developer of this part of code after time.

This tool tries to help. It is inspired by JavaScript and their API to manipulate objects.

# Install
`go get -u github.com/the-go-tool/object`

# Usage
If you familiar with JavaScript, you already familiar with this tool's methods.
Most of these methods constructed to work in chain-mode.

A bit of code to start:
```go
package main

import (
	"encoding/json"
	"github.com/the-go-tool/object"
)

func main() {
	source := []byte(`{"a":{"b":"c","d":-500.5},"e":[3, 2, 1]}`)
	var data interface{}
	_ = json.Unmarshal(source, &data)
	
	obj := object.New(data)
	
	obj.Get("a").Get("b").String() // "c"
	obj.GetIndex(1).Json() // "[3, 2, 3]" - getting keys by index and json marshaling
	obj.Get("e").GetIndex(0).String() // "3" - auto-convert if possible
	obj.Get("e").Get("0").Int() // 3 - same as above alternative
	
	obj.Get("not-exists").Get("d").IsExists() // false
	obj.Get("not-exists").Get("d").IsNull() // false - because it's not exists
	obj.Get("not-exists").Get("d").IsEmpty() // false - because it's not exists
	obj.Get("not-exists").Get("d").String() // "" - empty string, no error
	
	obj.Get("a").Get("d").IsValid() // true - exists & not null & not empty
	obj.Get("a").Get("d").IsNumber() // true - int, uint, int64, and etc allowed
	obj.Get("a").Get("d").Float64() // -500.5
	obj.Get("a").Get("d").Uint8() // 12
	
	obj.Get("a").Keys() // ["b", "d"]
	obj.Get("a").Get("b").Values() // [Object("c"), Object(-500.5)]
	obj.Get("a").Get("b").Entries() // [{Key: Object("b"), Value: Object("c")}, ...]
	
	obj.GetPath("e[1]").Int() // 2 - JavaScript-like syntax
}
```
