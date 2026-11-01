package component

import "fmt"

func Truncate(
	s string,
	n int,
) string {
	if len(s) <= n {
		return s
	}

	return fmt.Sprintf("%s...", s[:n-3])
}
