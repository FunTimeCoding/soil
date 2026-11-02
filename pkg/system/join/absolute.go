package join

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
	"strings"
)

func Absolute(parts ...string) string {
	parts = slice.Trim(slice.StripEmpty(parts), stringsConstant.Slash)

	if len(parts) == 0 {
		return stringsConstant.Slash
	}

	result := path.Join(parts...)

	if !strings.HasPrefix(result, stringsConstant.Slash) {
		result = key_value.Empty(stringsConstant.Slash, result)
	}

	return result
}
