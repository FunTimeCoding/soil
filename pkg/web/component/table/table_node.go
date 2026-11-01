package table

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (t *Table) tableNode() gomponents.Node {
	headers := make([]gomponents.Node, len(t.columns))

	for i, c := range t.columns {
		headers[i] = html.Th(
			gomponents.If(t.numeric[i], html.Class(constant.NumericClass)),
			gomponents.Text(c),
		)
	}

	return html.Table(
		gomponents.If(t.class != "", html.Class(t.class)),
		gomponents.If(!t.headerless, html.THead(html.Tr(headers...))),
		html.TBody(
			gomponents.If(t.bodyIdentifier != "", html.ID(t.bodyIdentifier)),
			gomponents.Group(t.rows),
		),
	)
}
