package unexpected

import "errors"

func Is(e error) bool {
	var target *UnexpectedError

	return errors.As(e, &target)
}
