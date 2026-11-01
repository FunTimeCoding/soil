package table

import "maragu.dev/gomponents"

func (t *Table) Rows(rows []gomponents.Node) *Table {
	t.rows = rows

	return t
}
