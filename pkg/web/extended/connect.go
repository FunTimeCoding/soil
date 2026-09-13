package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Connect(locator string) gomponents.Node {
	return gomponents.Attr(constant.ServerSideConnect, locator)
}
