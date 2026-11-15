package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func underTestHome(path string) bool {
	result := false

	for _, s := range strings.Split(path, "/") {
		if s == constant.TestdataDirectory {
			return false
		}

		if isTestHome(s) {
			result = true
		}
	}

	return result
}
