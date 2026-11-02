package build

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func Package(mainPath string) string {
	return join.Empty(
		soil.CurrentDirectory,
		string(filepath.Separator),
		filepath.Dir(mainPath),
	)
}
