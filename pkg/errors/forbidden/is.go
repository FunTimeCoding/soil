package forbidden

import "errors"

func Is(e error) bool {
	var target *ForbiddenError

	return errors.As(e, &target)
}
