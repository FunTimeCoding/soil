package unreachable

import "fmt"

func Format(
	format string,
	arguments ...any,
) *UnreachableError {
	return &UnreachableError{Message: fmt.Sprintf(format, arguments...)}
}
