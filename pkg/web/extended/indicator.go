package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Indicator(selector string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedIndicator, selector)
}
