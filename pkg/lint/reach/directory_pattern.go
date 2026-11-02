package reach

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func directoryPattern(relative string) string {
	if relative == soil.CurrentDirectory {
		return relative
	}

	return join.Empty(
		soil.CurrentDirectory,
		stringsConstant.Slash,
		filepath.ToSlash(relative),
	)
}
