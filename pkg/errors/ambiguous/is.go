package ambiguous

import "errors"

func Is(e error) bool {
	var target *AmbiguousError

	return errors.As(e, &target)
}
