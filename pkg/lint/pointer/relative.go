package pointer

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path"
	"strings"
)

func Relative(
	source string,
	s string,
) (string, bool) {
	result := path.Join(path.Dir(source), Normalize(s))

	return result, !strings.HasPrefix(result, constant.ParentDirectory)
}
