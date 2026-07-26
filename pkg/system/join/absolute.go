package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
	"strings"
)

func Absolute(parts ...string) string {
	parts = slice.Trim(slice.StripEmpty(parts), constant.Slash)

	if len(parts) == 0 {
		return constant.Slash
	}

	result := path.Join(parts...)

	if !strings.HasPrefix(result, constant.Slash) {
		result = key_value.Empty(constant.Slash, result)
	}

	return result
}
