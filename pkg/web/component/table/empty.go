package table

func (t *Table) Empty(message string) *Table {
	t.empty = message

	return t
}
