package entity_errors

import "errors"

var (
	ErrInvalidCollection = errors.New("invalid collection")
	ErrCollectionExists  = errors.New("collection already exists")
)
