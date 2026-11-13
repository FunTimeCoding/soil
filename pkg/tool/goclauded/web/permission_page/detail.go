package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"maragu.dev/gomponents"
)

func Detail(w *permission.Row) gomponents.Node {
	return opening(w).Detail(len(cells(w)), commands(w)).Response()
}
