package unavailable

import "errors"

func Is(e error) bool {
	var target *UnavailableError

	return errors.As(e, &target)
}
