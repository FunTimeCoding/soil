package refusal

import "maps"

func (t *Table) Rows() map[int]string {
	return maps.Clone(t.rows)
}
