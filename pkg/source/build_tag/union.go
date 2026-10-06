package build_tag

import (
	"maps"
	"slices"
)

func Union(files map[string][]string) []string {
	tags := make(map[string]bool)

	for _, t := range files {
		for _, tag := range t {
			tags[tag] = true
		}
	}

	return slices.Sorted(maps.Keys(tags))
}
