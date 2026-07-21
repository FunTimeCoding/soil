package scan

import "strings"

func isTestHomePath(path string) bool {
	segments := strings.Split(path, "/")
	limit := 3

	if len(segments) > 2 && segments[0] == "pkg" && segments[1] == "tool" {
		limit = 4
	}

	for i, s := range segments {
		if s == "testdata" {
			return true
		}

		if s != "unit_test" && s != "integration_test" {
			continue
		}

		if i+1 <= limit {
			return true
		}
	}

	return false
}
