package imports

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"slices"
)

func isBanned(name string) bool {
	return segment.Check(name, false, false) != nil ||
		slices.Contains(lintConstant.PackageBlocklist, name)
}
