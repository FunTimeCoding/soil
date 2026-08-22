package conflict

import "fmt"

func Format(
	format string,
	arguments ...any,
) *ConflictError {
	return &ConflictError{Message: fmt.Sprintf(format, arguments...)}
}
