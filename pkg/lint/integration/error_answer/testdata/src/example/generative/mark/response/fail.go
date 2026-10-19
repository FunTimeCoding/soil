package response

import "fmt"

func Fail(
	format string,
	a ...any,
) (string, error) {
	return fmt.Sprintf(format, a...), nil
}
