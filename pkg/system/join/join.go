package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
	"strings"
)

func Join(parts ...string) string {
	result := slice.StripEmpty(parts)
	absolute := strings.HasPrefix(result[0], constant.Slash)
	result = slice.Trim(result, constant.Slash)

	if len(result) == 0 {
		return constant.Slash
	}

	if absolute {
		result = slice.Prepend(result, constant.Slash)
	}

	return path.Join(result...)
}
