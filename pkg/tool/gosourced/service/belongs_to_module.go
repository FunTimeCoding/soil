package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func belongsToModule(
	packagePath string,
	modulePath string,
) bool {
	return packagePath == modulePath ||
		strings.HasPrefix(packagePath, join.Empty(modulePath, "/"))
}
