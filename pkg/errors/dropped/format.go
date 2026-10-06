package dropped

import "fmt"

func Format(
	format string,
	arguments ...any,
) *DroppedError {
	return &DroppedError{Message: fmt.Sprintf(format, arguments...)}
}
