package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
)

func Relative(parts ...string) string {
	return path.Join(slice.Trim(slice.StripEmpty(parts), constant.Slash)...)
}
