package not_selected

import "fmt"

func Format(
	format string,
	arguments ...any,
) *NotSelectedError {
	return &NotSelectedError{Message: fmt.Sprintf(format, arguments...)}
}
