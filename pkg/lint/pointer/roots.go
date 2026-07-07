package pointer

import (
	"sort"
	"strings"
)

func Roots(paths []string) []string {
	seen := map[string]bool{}

	for _, p := range paths {
		if root, _, found := strings.Cut(p, "/"); found {
			seen[root] = true
		}
	}

	var result []string

	for root := range seen {
		result = append(result, root)
	}

	sort.Strings(result)

	return result
}
