package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func isConstantRoot(directory string) bool {
	segments := strings.Split(directory, "/")

	if segments[0] != constant.PackageDirectory {
		return false
	}

	switch len(segments) {
	case 2:
		return true
	case 3:
		return segments[1] != constant.ToolDirectory
	case 4:
		return segments[1] == constant.ToolDirectory
	}

	return false
}
