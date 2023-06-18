package custom_error

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Err        error  // logging
	Code       int    // http status code
	Message    string // informing client
	ChildError *CustomError
}

func (e *CustomError) Error() string {
	if e.ChildError != nil {
		return fmt.Errorf("%w: %s", e.Err, e.ChildError.Error()).Error()
	}
	return fmt.Errorf("%w: %s", e.Err, e.Message).Error()
}

func New(err error, typeCode int, message string) *CustomError {
	return &CustomError{
		Err:     err,
		Code:    typeCode,
		Message: message,
	}
}

func (e *CustomError) Wrap(err string) *CustomError {
	return &CustomError{
		Err:        errors.New(err),
		Code:       e.Code,
		Message:    e.Message,
		ChildError: e,
	}
}
