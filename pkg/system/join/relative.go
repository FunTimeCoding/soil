package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"path"
)

func Relative(parts ...string) string {
	return path.Join(slice.Trim(slice.StripEmpty(parts), separator.Slash)...)
}
