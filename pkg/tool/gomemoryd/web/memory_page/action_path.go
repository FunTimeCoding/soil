package memory_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func actionPath(
	identifier int64,
	action string,
) string {
	return join.Empty(Path(identifier), stringsConstant.Slash, action)
}
