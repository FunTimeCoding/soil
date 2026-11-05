package refusal

import "maps"

func (t *Table) With(
	status int,
	class string,
) *Table {
	result := &Table{rows: maps.Clone(t.rows)}
	result.rows[status] = class

	return result
}
