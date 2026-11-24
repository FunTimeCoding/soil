package remediable

import "fmt"

func Format(
	remedy string,
	format string,
	arguments ...any,
) *RemediableError {
	return &RemediableError{
		Name:    remedy,
		Message: fmt.Sprintf(format, arguments...),
	}
}
