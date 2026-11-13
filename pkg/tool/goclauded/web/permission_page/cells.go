package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func cells(w *permission.Row) []gomponents.Node {
	return []gomponents.Node{
		html.Td(gomponents.Text(w.Blocker)),
		html.Td(gomponents.Textf("%d", w.Prompts)),
		html.Td(gomponents.Textf("%d", w.Sessions)),
		html.Td(gomponents.Text(reason(w.Reason))),
		layout.TimeCell(w.Last),
	}
}
