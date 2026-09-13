package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Confirm(message string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedConfirm, message)
}
