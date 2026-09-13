package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Get(locator string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedGet, locator)
}
