package not_found

import "fmt"

func Format(
	format string,
	arguments ...any,
) *NotFoundError {
	return &NotFoundError{Message: fmt.Sprintf(format, arguments...)}
}
