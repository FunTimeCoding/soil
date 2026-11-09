package symbol

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func inWorkspace(
	modules []string,
	packagePath string,
) bool {
	for _, m := range modules {
		if packagePath == m ||
			strings.HasPrefix(packagePath, join.Empty(m, "/")) {
			return true
		}
	}

	return false
}
