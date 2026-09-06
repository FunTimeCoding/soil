package scan

import "strings"

func coveredRoute(
	patterns []string,
	route string,
) bool {
	for _, pattern := range patterns {
		if pattern == route {
			return true
		}

		if strings.HasSuffix(pattern, "/") &&
			strings.HasPrefix(route, pattern) {
			return true
		}
	}

	return false
}
