package timeout

import "fmt"

func Format(
	format string,
	arguments ...any,
) *TimeoutError {
	return &TimeoutError{Message: fmt.Sprintf(format, arguments...)}
}
