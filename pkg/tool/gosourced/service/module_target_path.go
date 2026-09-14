package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func moduleTargetPath(
	packagePath string,
	modulePath string,
	newModulePath string,
) string {
	return join.Empty(
		newModulePath,
		strings.TrimPrefix(packagePath, modulePath),
	)
}
