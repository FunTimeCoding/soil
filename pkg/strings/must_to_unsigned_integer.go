package strings

import (
	"strconv"
	"strings"
)

func MustToUnsignedInteger(i string) uint {
	result, e := strconv.ParseUint(strings.TrimSpace(i), 10, 64)

	if e != nil {
		panic(e)
	}

	return uint(result)
}
