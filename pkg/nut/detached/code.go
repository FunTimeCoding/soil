package detached

import "errors"

func Code(e error) string {
	var target *DetachedError

	if !errors.As(e, &target) {
		return ""
	}

	return target.Code
}
