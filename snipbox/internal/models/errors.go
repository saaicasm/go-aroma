package models

import "errors"

var (
	ErrNoRecord           = errors.New("models : no matching record found")
	ErrInvalidCredentials = errors.New("users: Invalid credentials")
	ErrDuplicateEmail     = errors.New("users: Duplicate Email address")
)
