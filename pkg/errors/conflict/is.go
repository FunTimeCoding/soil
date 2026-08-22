package conflict

import "errors"

func Is(e error) bool {
	var target *ConflictError

	return errors.As(e, &target)
}
