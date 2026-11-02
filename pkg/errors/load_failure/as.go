package load_failure

import "errors"

func As(e error) (*LoadFailureError, bool) {
	var target *LoadFailureError

	return target, errors.As(e, &target)
}
