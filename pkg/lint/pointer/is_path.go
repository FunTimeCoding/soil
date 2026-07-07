package pointer

import "strings"

func IsPath(s string) bool {
	if s == "" || strings.ContainsAny(s, " \t") {
		return false
	}

	return strings.Contains(s, "/")
}
