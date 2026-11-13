package table

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (t *Table) tableNode() gomponents.Node {
	headers := make([]gomponents.Node, len(t.columns))

	for i, c := range t.columns {
		headers[i] = html.Th(gomponents.Text(c))
	}

	class := t.tableClass()

	return html.Table(
		gomponents.If(class != "", html.Class(class)),
		gomponents.If(!t.headerless, html.THead(html.Tr(headers...))),
		html.TBody(
			gomponents.If(t.bodyIdentifier != "", html.ID(t.bodyIdentifier)),
			gomponents.Group(t.rows),
		),
	)
}
