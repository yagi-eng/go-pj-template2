package myerror

import "github.com/pkg/errors"

type MyError struct {
	Code errorCode
	Err  error
}

func (e *MyError) Error() string {
	return e.Err.Error()
}

func New(code errorCode, err error) *MyError {
	return &MyError{
		Code: code,
		Err:  errors.WithStack(err),
	}
}
