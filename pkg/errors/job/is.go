package job

import "errors"

func Is(e error) bool {
	var target *JobError

	return errors.As(e, &target)
}
