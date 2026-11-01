package component

import "fmt"

func PerMinute(
	total int,
	minutes float64,
) string {
	if minutes <= 0 {
		return "-"
	}

	return fmt.Sprintf("%.2f", float64(total)/minutes)
}
