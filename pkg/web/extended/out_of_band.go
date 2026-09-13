package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func OutOfBand(value string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedOutOfBand, value)
}
