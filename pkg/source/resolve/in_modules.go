package resolve

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func inModules(
	path string,
	modules []string,
) bool {
	for _, m := range modules {
		if path == m || strings.HasPrefix(path, join.Empty(m, constant.Slash)) {
			return true
		}
	}

	return false
}
