package service

import "errors"

var (
	ErrValidation = errors.New("invalid")
	ErrNotFound   = errors.New("not found")
	ErrDuplicate  = errors.New("duplicate")
	ErrInternal   = errors.New("internal error")
)
