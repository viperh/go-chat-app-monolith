package users

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExist       = errors.New("user already exist")
	ErrInvalidPassword = errors.New("invalid password")
)
