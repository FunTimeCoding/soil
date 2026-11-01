package table

func (t *Table) Identifier(identifier string) *Table {
	t.identifier = identifier

	return t
}
