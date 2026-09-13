package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Post(locator string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedPost, locator)
}
