package strings

import (
	"strconv"
	"strings"
)

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}
