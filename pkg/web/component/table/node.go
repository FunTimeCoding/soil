package table

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (t *Table) Node() gomponents.Node {
	content := t.tableNode()

	if t.empty != "" && len(t.rows) == 0 {
		content = html.P(
			html.Class(constant.AbsentClass),
			gomponents.Text(t.empty),
		)
	}

	if t.identifier == "" {
		return content
	}

	return html.Div(html.ID(t.identifier), content)
}
