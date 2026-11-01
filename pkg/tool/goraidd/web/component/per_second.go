package component

import "fmt"

func PerSecond(
	total int,
	seconds float64,
	decimals int,
) string {
	if seconds <= 0 {
		return "-"
	}

	format := fmt.Sprintf("%%.%df", decimals)

	return fmt.Sprintf(format, float64(total)/seconds)
}
