package resolve

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"slices"
)

func WithMainModule(patterns []string) []string {
	if CoversMainModule(patterns) {
		return patterns
	}

	return append(slices.Clone(patterns), goModConstant.AllPackages)
}
