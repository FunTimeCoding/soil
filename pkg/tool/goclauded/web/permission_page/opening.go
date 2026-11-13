package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
)

func opening(w *permission.Row) *opening_row.OpeningRow {
	return opening_row.New(
		component.RowMark(constant.PermissionRowPrefix, w.Blocker),
		component.RowPath(
			constant.PermissionsPath,
			constant.OpenAction,
			constant.PermissionBlocker,
			w.Blocker,
		),
		component.RowPath(
			constant.PermissionsPath,
			constant.CloseAction,
			constant.PermissionBlocker,
			w.Blocker,
		),
	).Cells(cells(w)...)
}
