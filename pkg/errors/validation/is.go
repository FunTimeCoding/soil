package validation

import "errors"

func Is(e error) bool {
	var target *Detail

	return errors.As(e, &target)
}
