package errors

import "context"

func Canceled(e error) bool {
	return Is(e, context.Canceled)
}
