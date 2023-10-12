package graphapi

import (
	"errors"
	"fmt"
)

// ErrInternalServerError is returned when an internal error occurs.
var ErrInternalServerError = errors.New("internal server error")

// ErrInvalidID is returned when an invalid ID is provided.
type ErrInvalidID struct {
	field string
	err   error
}

// Error implements the error interface.
func (e *ErrInvalidID) Error() string {
	return fmt.Sprintf("%v, field: %s", e.err, e.field)
}
