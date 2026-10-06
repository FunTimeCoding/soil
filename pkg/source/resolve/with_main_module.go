package resolve

import "slices"

func WithMainModule(patterns []string) []string {
	if CoversMainModule(patterns) {
		return patterns
	}

	return append(slices.Clone(patterns), "./...")
}
