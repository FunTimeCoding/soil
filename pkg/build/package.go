package build

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func Package(mainPath string) string {
	return join.Empty(
		constant.CurrentDirectory,
		string(filepath.Separator),
		filepath.Dir(mainPath),
	)
}
