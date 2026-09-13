package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Swap(style string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedSwap, style)
}
