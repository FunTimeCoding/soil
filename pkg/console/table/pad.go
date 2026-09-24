package table

import "fmt"

func pad(
	s string,
	width int,
	right bool,
) string {
	if len(s) >= width {
		return s
	}

	if right {
		return fmt.Sprintf("%*s", width, s)
	}

	return fmt.Sprintf("%-*s", width, s)
}
