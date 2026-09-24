package table

func (t *Table) Right(indexes ...int) *Table {
	for _, i := range indexes {
		if i < len(t.columns) {
			t.columns[i].right = true
		}
	}

	return t
}
