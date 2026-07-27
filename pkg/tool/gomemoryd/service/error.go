package service

import "errors"

var (
	ErrorAlwaysLoad     = errors.New("failed to load always memories")
	ErrorRelevantSearch = errors.New("failed to search relevant memories")
	ErrorMemoryList     = errors.New("failed to list memories")
	ErrorReservedScope  = errors.New("scope name is reserved")
)
