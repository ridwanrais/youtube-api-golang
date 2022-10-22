package shared

import (
	"errors"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func CreateError(code int, message string) error {
	return &RequestError{
		StatusCode: code,
		Err:        errors.New(message),
	}
}
