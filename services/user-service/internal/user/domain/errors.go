package domain

import "errors"


var ErrUserNotFound = errors.New(
	"user not found",
)

var ErrInvalidUserID = errors.New(
	"invalid user id",
)
