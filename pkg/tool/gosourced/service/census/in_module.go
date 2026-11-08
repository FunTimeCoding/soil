package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func inModule(
	module string,
	path string,
) bool {
	return path == module || strings.HasPrefix(path, join.Empty(module, "/"))
}
