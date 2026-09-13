package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Trigger(when string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedTrigger, when)
}
