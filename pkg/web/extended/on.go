package extended

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func On(
	event string,
	script string,
) gomponents.Node {
	return gomponents.Attr(join.Empty(constant.ExtendedOnPrefix, event), script)
}
