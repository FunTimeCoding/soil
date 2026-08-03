package not_found

import "errors"

func Is(e error) bool {
	var target *NotFoundError

	return errors.As(e, &target)
}
