package strings

import (
	"strconv"
	"strings"
)

func ToUnsignedInteger(
	i string,
	fallback uint,
) uint {
	if result, e := strconv.ParseUint(strings.TrimSpace(i), 10, 64); e == nil {
		return uint(result)
	}

	return fallback
}
