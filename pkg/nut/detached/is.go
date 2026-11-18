package detached

import "errors"

func Is(e error) bool {
	var target *DetachedError

	return errors.As(e, &target)
}
