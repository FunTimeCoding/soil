package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Include(selector string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedInclude, selector)
}
