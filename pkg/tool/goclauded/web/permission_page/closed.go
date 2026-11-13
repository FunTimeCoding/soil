package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"maragu.dev/gomponents"
)

func Closed(w *permission.Row) gomponents.Node {
	return opening(w).Response()
}
