package table

func (t *Table) BodyIdentifier(identifier string) *Table {
	t.bodyIdentifier = identifier

	return t
}
