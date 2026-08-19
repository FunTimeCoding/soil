package web

import "slices"

func matchesFilter(
	documentMetadata map[string][]string,
	filter map[string]string,
) bool {
	for key, value := range filter {
		if !slices.Contains(documentMetadata[key], value) {
			return false
		}
	}

	return true
}
