package scan

import "github.com/funtimecoding/soil/pkg/tool/goaudit/constant"

func isTestdataPath(segments []string) bool {
	for _, s := range segments {
		if s == constant.TestdataDirectory {
			return true
		}
	}

	return false
}
