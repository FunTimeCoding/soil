package caller_left

import "errors"

func Is(e error) bool {
	var target *CallerLeft

	return errors.As(e, &target)
}
