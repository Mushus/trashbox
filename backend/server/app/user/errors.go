package user

import (
	"golang.org/x/xerrors"
)

var (
	// ErrUserNotFound occurs when the user is not found
	ErrUserNotFound = xerrors.New("user not found")
	// ErrInvalidLoginOrPassword occurs when the user input invalid login or password
	ErrInvalidLoginOrPassword = xerrors.New("invalid user or password")
)
