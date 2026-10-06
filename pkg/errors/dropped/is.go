package dropped

import "errors"

func Is(e error) bool {
	var target *DroppedError

	return errors.As(e, &target)
}
