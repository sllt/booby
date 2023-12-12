package pubsub

import "errors"

var (
	// ErrInvalidPassword .
	ErrInvalidPassword = errors.New("invalid password")

	// ErrInvalidTopicEmpty .
	ErrInvalidTopicEmpty = errors.New("invalid topic, should not be \"\"")

	// ErrInvalidTopicBytes .
	ErrInvalidTopicBytes = errors.New("invalid topic bytes, should be more than 10 bytes")

	// ErrInvalidTopicNameLength .
	ErrInvalidTopicNameLength = errors.New("invalid topic name length, should not be more than 1024")
)
