package models

import (
	"errors"
)

// Common app errors.
//
// If error is handled in the server, the corresponding HTTP status code is indicated.
var (
	// ErrAccessDenied is returned when a user does not have permission to access a resource.
	//
	// Server responds with 404 Not Found.
	ErrAccessDenied = errors.New("access denied")

	// ErrDuplicate is returned when a resource already exists.
	ErrDuplicate = errors.New("duplicate")

	// ErrInvalidAuth is returned when authentication fails.
	//
	// Server responds with 401 Unauthorized.
	ErrInvalidAuth = errors.New("invalid auth")

	// ErrNotFound is returned when a resource is not found.
	//
	// Server responds with 404 Not Found.
	ErrNotFound = errors.New("not found")

	ErrPanic = errors.New("panic")
)
