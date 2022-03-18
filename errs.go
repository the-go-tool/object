package object

import (
	"errors"
	"fmt"
)

const (
	ErrorObjectNotExists = "object isn't exists"
	ErrorTypeNotSupport  = "type isn't supporting"
	ErrorFieldNotFound   = "field name not found"
	ErrorIndexParse      = "index can't be parsed"
	ErrorIndexRange      = "index out of range"
	ErrorDataParse       = "data can't be parsed"
)

// Error - objects manipulation error
type Error struct {
	err error
}

func newError(text string) *Error {
	return &Error{
		err: errors.New(text),
	}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%v", e.err)
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *Error) Is(target error) bool {
	return e.err == target
}
