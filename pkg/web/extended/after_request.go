package extended

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func AfterRequest(script string) gomponents.Node {
	return gomponents.Attr(constant.ExtendedAfterRequest, script)
}
