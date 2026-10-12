package imports

import (
	"strings"
	"unicode"
)

func IsVersion(name string) bool {
	digits, okay := strings.CutPrefix(name, "v")

	if !okay || digits == "" {
		return false
	}

	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
