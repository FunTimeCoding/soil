package imports

import "github.com/funtimecoding/soil/pkg/source/constant"

func ConstantAlias(
	importPath string,
	modulePath string,
) string {
	s := Subsystem(importPath, modulePath)

	if s == "" || IsLayoutRoot(importPath) {
		return s
	}

	return Fallback(s, constant.ConstantPackage)
}
