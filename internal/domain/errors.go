package domain

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("resource not found")
	ErrConflict     = errors.New("resource already exists")
	ErrInternal     = errors.New("internal server error")
	ErrInvalidDate  = errors.New("invalid date format, expected YYYY-MM-DD")
)
