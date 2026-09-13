package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Extension(name string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedExtension, name)
}
