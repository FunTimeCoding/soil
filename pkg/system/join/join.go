package join

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
	"strings"
)

func Join(parts ...string) string {
	result := slice.StripEmpty(parts)
	absolute := strings.HasPrefix(result[0], stringsConstant.Slash)
	result = slice.Trim(result, stringsConstant.Slash)

	if len(result) == 0 {
		return stringsConstant.Slash
	}

	if absolute {
		result = slice.Prepend(result, stringsConstant.Slash)
	}

	return path.Join(result...)
}
