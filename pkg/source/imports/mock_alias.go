package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
)

func MockAlias(
	importPath string,
	modulePath string,
) string {
	for directory := importPath; directory != "."; directory = path.Dir(directory) {
		s := Subsystem(directory, modulePath)

		if s == "" {
			return ""
		}

		if !constant.MockOwnerSkips[s] {
			return join.Empty(lowerCamel(s), constant.MockSuffix)
		}
	}

	return ""
}
