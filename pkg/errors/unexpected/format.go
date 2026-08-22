package unexpected

import "fmt"

func Format(
	format string,
	arguments ...any,
) *UnexpectedError {
	return &UnexpectedError{Message: fmt.Sprintf(format, arguments...)}
}
