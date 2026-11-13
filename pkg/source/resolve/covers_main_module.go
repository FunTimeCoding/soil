package resolve

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"slices"
)

func CoversMainModule(patterns []string) bool {
	return slices.Contains(patterns, goModConstant.AllPackages)
}
