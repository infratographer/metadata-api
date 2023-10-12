package graphapi

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError is returned when an internal error occurs.
	ErrInternalServerError = errors.New("internal server error")

	// ErrInvalidJSON is returned when invalid json is provided.
	ErrInvalidJSON = errors.New("invalid json data")
)

// ErrInvalidID is returned when an invalid ID is provided.
type ErrInvalidID struct {
	field string
	err   error
}

// Error implements the error interface.
func (e *ErrInvalidID) Error() string {
	return fmt.Sprintf("%v, field: %s", e.err, e.field)
}
