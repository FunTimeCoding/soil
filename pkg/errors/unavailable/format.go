package unavailable

import "fmt"

func Format(
	format string,
	arguments ...any,
) *UnavailableError {
	return &UnavailableError{Message: fmt.Sprintf(format, arguments...)}
}
