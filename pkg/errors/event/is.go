package event

import "errors"

func Is(e error) bool {
	var target *EventError

	return errors.As(e, &target)
}
