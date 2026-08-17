package tracker

import "strings"

func follow(
	name string,
	prefixes []string,
) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}

	return false
}
