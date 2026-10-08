package search_index

import (
	"maps"
	"slices"
)

func stringValues(v any) []string {
	switch t := v.(type) {
	case string:
		return []string{t}
	case []any:
		var result []string

		for _, item := range t {
			result = append(result, stringValues(item)...)
		}

		return result
	case map[string]any:
		var result []string

		for _, key := range slices.Sorted(maps.Keys(t)) {
			result = append(result, stringValues(t[key])...)
		}

		return result
	}

	return nil
}
