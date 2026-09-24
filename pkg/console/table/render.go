package table

import "strings"

func (t *Table) Render() string {
	var b strings.Builder
	t.line(&b, t.headers())

	for _, row := range t.rows {
		t.line(&b, row)
	}

	return b.String()
}
