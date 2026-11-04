package reported

import "errors"

func Is(e error) bool {
	var target *ReportedError

	return errors.As(e, &target)
}
