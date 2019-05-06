package repository

import (
	"golang.org/x/xerrors"
)

var (
	// ErrUserNotFound occurs when repository find a user but user not found it
	ErrUserNotFound = xerrors.New("user not found")
	// ErrDocumentNotFound is a error thrown when the document does not exist
	ErrDocumentNotFound = xerrors.New("document not found")
)
