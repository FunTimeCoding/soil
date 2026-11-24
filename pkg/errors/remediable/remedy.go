package remediable

import "errors"

func Remedy(e error) string {
	var target *RemediableError

	if !errors.As(e, &target) {
		return ""
	}

	return target.Name
}
