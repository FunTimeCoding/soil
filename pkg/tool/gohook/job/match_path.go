package job

import (
	"github.com/bmatcuk/doublestar/v4"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func matchPath(
	pattern string,
	path string,
) bool {
	if strings.HasSuffix(pattern, constant.Slash) {
		return strings.HasPrefix(path, pattern)
	}

	if strings.ContainsAny(pattern, "*?[") {
		result, e := doublestar.Match(pattern, path)
		errors.PanicOnError(e)

		return result
	}

	return pattern == path
}
