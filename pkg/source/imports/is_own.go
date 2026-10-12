package imports

import (
	"path"
	"strings"
)

func IsOwn(
	packagePath string,
	importPath string,
) bool {
	if IsLayoutRoot(importPath) {
		return packagePath == importPath
	}

	rest, okay := strings.CutPrefix(packagePath, path.Dir(importPath))

	return okay && (rest == "" || strings.HasPrefix(rest, "/"))
}
