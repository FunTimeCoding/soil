package event

import "strings"

func hasMemoryKind(kinds []string) bool {
	for _, k := range kinds {
		if strings.HasPrefix(k, "memory_") {
			return true
		}
	}

	return false
}
