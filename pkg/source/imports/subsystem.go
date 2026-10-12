package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"go/token"
	"path"
	"strings"
)

func Subsystem(
	importPath string,
	modulePath string,
) string {
	if importPath == modulePath {
		return ""
	}

	inModule := modulePath != "" && strings.HasPrefix(importPath, modulePath)

	for directory := path.Dir(importPath); directory != "."; directory = path.Dir(directory) {
		name := path.Base(directory)
		last := inModule && directory == modulePath

		if constant.LayoutRoots[name] {
			if !inModule {
				return ""
			}

			name = path.Base(modulePath)
			last = true
		}

		if token.IsIdentifier(name) && !isBanned(name) {
			return name
		}

		if last {
			return ""
		}
	}

	return ""
}
