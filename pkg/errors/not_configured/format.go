package not_configured

import "fmt"

func Format(
	format string,
	arguments ...any,
) *NotConfiguredError {
	return &NotConfiguredError{Message: fmt.Sprintf(format, arguments...)}
}
