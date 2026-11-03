package imports

import (
	"strings"
	"unicode"
)

func IsNumberedFallback(
	local string,
	declaredName string,
) bool {
	digits, okay := strings.CutPrefix(local, declaredName)

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
