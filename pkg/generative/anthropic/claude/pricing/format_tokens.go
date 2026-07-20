package pricing

import (
	"fmt"
	"strconv"
)

func FormatTokens(count int) string {
	if count >= 1_000_000_000 {
		return fmt.Sprintf("%.1fB", float64(count)/1_000_000_000)
	}

	if count >= 1_000_000 {
		return fmt.Sprintf("%.1fM", float64(count)/1_000_000)
	}

	if count >= 1_000 {
		return fmt.Sprintf("%.1fK", float64(count)/1_000)
	}

	return strconv.Itoa(count)
}
