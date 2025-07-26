package models

import "errors"

// Common app errors.
var (
	ErrDuplicate   = errors.New("duplicate")
	ErrInvalidAuth = errors.New("invalid auth")
	ErrNotFound    = errors.New("not found")
	ErrPanic       = errors.New("panic")
)
