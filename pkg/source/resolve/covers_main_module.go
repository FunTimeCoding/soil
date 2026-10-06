package resolve

import "slices"

func CoversMainModule(patterns []string) bool {
	return slices.Contains(patterns, "./...")
}
