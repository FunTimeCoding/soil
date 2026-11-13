package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"maragu.dev/gomponents"
)

func Removal(blocker string) gomponents.Node {
	return opening_row.New(
		component.RowMark(constant.PermissionRowPrefix, blocker),
		"",
		"",
	).Removal()
}
