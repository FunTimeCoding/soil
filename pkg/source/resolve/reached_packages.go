package resolve

import (
	"golang.org/x/tools/go/packages"
	"sort"
)

func ReachedPackages(
	loaded []*packages.Package,
	modules []string,
) []string {
	if len(modules) == 0 {
		return nil
	}

	seen := make(map[string]bool)
	reached := make(map[string]bool)

	for _, p := range loaded {
		if p.Types != nil {
			collectReached(p.Types, modules, seen, reached)
		}
	}

	result := make([]string, 0, len(reached))

	for path := range reached {
		result = append(result, path)
	}

	sort.Strings(result)

	return result
}
