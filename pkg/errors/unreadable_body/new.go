package unreadable_body

import "fmt"

func New(
	wrapped error,
	format string,
	arguments ...any,
) *UnreadableBodyError {
	return &UnreadableBodyError{
		Message: fmt.Sprintf(format, arguments...),
		Wrapped: wrapped,
	}
}
