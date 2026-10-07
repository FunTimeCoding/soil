package table

func (t *Table) Add(values ...string) {
	for i, v := range values {
		if i < len(t.columns) && len(v) > t.columns[i].Width {
			t.columns[i].Width = len(v)
		}
	}

	t.rows = append(t.rows, values)
}
