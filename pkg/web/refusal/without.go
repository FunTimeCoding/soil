package refusal

import "maps"

func (t *Table) Without(status int) *Table {
	result := &Table{rows: maps.Clone(t.rows)}
	delete(result.rows, status)

	return result
}
