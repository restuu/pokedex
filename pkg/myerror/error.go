package myerror

import "errors"

var (
	ErrDataNotFound  = errors.New("data not found")
	ErrInvalidFormat = errors.New("invalid format")
)
