package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"maragu.dev/gomponents"
)

func Closed(g *upgrade_group.Group) gomponents.Node {
	return opening(g).Response()
}
