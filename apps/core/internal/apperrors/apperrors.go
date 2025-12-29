package apperrors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidIdFormat   = errors.New("invalid id format")
	ErrUnexpected        = errors.New("unexpected error occurred")
	ErrUserAlreadyExists = errors.New("user already exists")
)
