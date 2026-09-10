package refusal

import "errors"

func Is(e error) bool {
	var target *Refusal

	return errors.As(e, &target)
}
