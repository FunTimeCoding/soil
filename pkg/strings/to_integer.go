package strings

import (
	"strconv"
	"strings"
)

func ToInteger(
	i string,
	fallback int,
) int {
	if result, e := strconv.ParseInt(strings.TrimSpace(i), 10, 32); e == nil {
		return int(result)
	}

	return fallback
}
