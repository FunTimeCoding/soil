package command

import "errors"

func Is(e error) bool {
	var target *CommandError

	return errors.As(e, &target)
}
