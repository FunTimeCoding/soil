package unreadable_body

import "errors"

func Is(e error) bool {
	var target *UnreadableBodyError

	return errors.As(e, &target)
}
