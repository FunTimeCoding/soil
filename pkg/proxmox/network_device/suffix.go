package network_device

import (
	"strconv"
	"strings"
	"unicode"
)

func suffix(v string) int {
	i := strings.IndexFunc(v, unicode.IsDigit)

	if i < 0 {
		return 0
	}

	result, e := strconv.Atoi(v[i:])

	if e != nil {
		return 0
	}

	return result
}
