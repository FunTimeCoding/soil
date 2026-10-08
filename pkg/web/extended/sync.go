package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Sync(strategy string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedSync, strategy)
}
