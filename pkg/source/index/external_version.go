package index

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"runtime"
	"strings"
)

func externalVersion(
	path string,
	requirements map[string]string,
) string {
	var best string

	for module := range requirements {
		if (path == module ||
			strings.HasPrefix(path, join.Empty(module, constant.Slash))) &&
			len(module) > len(best) {
			best = module
		}
	}

	if best == "" {
		return runtime.Version()
	}

	return requirements[best]
}
