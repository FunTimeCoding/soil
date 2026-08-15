package lint

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
	"strings"
)

func testdataPath(path string) bool {
	for _, s := range strings.Split(
		filepath.ToSlash(filepath.Dir(path)),
		"/",
	) {
		if s == constant.TestdataPath {
			return true
		}
	}

	return false
}
