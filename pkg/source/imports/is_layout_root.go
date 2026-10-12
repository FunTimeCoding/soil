package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"path"
)

func IsLayoutRoot(importPath string) bool {
	return constant.LayoutRoots[path.Base(path.Dir(importPath))]
}
