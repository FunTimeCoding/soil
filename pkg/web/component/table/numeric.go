package table

func (t *Table) Numeric(columns ...int) *Table {
	for _, c := range columns {
		t.numeric[c] = true
	}

	return t
}
