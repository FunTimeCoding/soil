package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func InScope(
	o *option.Lint,
	path string,
) bool {
	if len(o.Scopes) == 0 {
		return true
	}

	for _, s := range o.Scopes {
		if path == s ||
			strings.HasPrefix(path, join.Empty(s, stringsConstant.Slash)) {
			return true
		}
	}

	return false
}
