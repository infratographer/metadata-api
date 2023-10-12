package client

import (
	"errors"
	"fmt"
)

// ErrInvaliID returned when an invalid id is provided
type ErrInvalidID struct {
	field string
	err   error
}

// Error implements the error interface.
func (e *ErrInvalidID) Error() string {
	return fmt.Sprintf("%v, field: %s", e.err, e.field)
}

var (
	// ErrUnauthorized returned when the request is not authorized
	ErrUnauthorized = errors.New("client is unauthorized")

	// ErrStatusNamespaceNotfound returned when the metadata status namespace is not found
	ErrStatusNamespaceNotfound = errors.New("status_namespace not found")

	// ErrHTTPError returned when the http response is an error
	ErrHTTPError = errors.New("metadata api http error")
)
