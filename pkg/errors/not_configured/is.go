package not_configured

import "errors"

func Is(e error) bool {
	var target *NotConfiguredError

	return errors.As(e, &target)
}
