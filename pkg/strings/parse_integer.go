package strings

import (
	"strconv"
	"strings"
)

func ParseInteger(s string) (int, error) {
	result, e := strconv.ParseInt(strings.TrimSpace(s), 10, 32)

	if e != nil {
		return 0, e
	}

	return int(result), nil
}
