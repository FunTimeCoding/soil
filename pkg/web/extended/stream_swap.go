package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func StreamSwap(name string) gomponents.Node {
	return gomponents.Attr(constant.ServerSideSwap, name)
}
