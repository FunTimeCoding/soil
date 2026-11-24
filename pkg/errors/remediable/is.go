package remediable

import "errors"

func Is(e error) bool {
	var target *RemediableError

	return errors.As(e, &target)
}
