package unreachable

import "errors"

func Is(e error) bool {
	var target *UnreachableError

	return errors.As(e, &target)
}
