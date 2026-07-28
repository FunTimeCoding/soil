package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func isTestHomePath(path string) bool {
	segments := strings.Split(path, "/")
	limit := 3

	if len(segments) > 2 &&
		segments[0] == constant.PackageDirectory &&
		segments[1] == constant.ToolDirectory {
		limit = 4
	}

	for i, s := range segments {
		if s == constant.TestdataDirectory {
			return true
		}

		if !isTestHome(s) {
			continue
		}

		if i+1 <= limit {
			return true
		}
	}

	return false
}
