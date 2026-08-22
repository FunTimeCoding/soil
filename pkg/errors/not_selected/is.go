package not_selected

import "errors"

func Is(e error) bool {
	var target *NotSelectedError

	return errors.As(e, &target)
}
