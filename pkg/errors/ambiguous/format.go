package ambiguous

import "fmt"

func Format(
	format string,
	arguments ...any,
) *AmbiguousError {
	return &AmbiguousError{Message: fmt.Sprintf(format, arguments...)}
}
