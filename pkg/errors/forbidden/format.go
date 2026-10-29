package forbidden

import "fmt"

func Format(
	format string,
	arguments ...any,
) *ForbiddenError {
	return &ForbiddenError{Message: fmt.Sprintf(format, arguments...)}
}
