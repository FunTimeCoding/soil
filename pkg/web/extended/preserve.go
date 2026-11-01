package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Preserve() gomponents.Node {
	return gomponents.Attr(constant.ExtendedPreserve, "true")
}
