package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func row(w *permission.Row) gomponents.Node {
	if len(w.Commands) == 0 {
		return html.Tr(gomponents.Group(cells(w)))
	}

	return opening(w).Node()
}
