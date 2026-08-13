package strings

import (
	"strconv"
	"strings"
)

func ToFloat(
	f string,
	fallback float64,
) float64 {
	if result, e := strconv.ParseFloat(strings.TrimSpace(f), 64); e == nil {
		return result
	}

	return fallback
}
