package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Target(selector string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedTarget, selector)
}
