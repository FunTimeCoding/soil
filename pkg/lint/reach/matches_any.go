package reach

import "path/filepath"

func matchesAny(
	patterns []string,
	path string,
) bool {
	for _, p := range patterns {
		if found, e := filepath.Match(p, path); e == nil && found {
			return true
		}
	}

	return false
}
