package util

import "github.com/go-playground/validator/v10"

// CustomValidator custom struct validator
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validate interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
