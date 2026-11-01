package table

func (t *Table) Headerless() *Table {
	t.headerless = true

	return t
}
