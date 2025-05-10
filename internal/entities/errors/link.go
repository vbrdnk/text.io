package entity_errors

import "errors"

var (
	ErrInvalidUrl  = errors.New("invalid link: invalid url")
	ErrInvalidLink = errors.New("invalid link")
	ErrLinkExists  = errors.New("link already exists")
)
