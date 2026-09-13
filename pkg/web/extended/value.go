package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Value(notation string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedValue, notation)
}
