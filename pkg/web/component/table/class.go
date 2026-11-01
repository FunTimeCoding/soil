package table

func (t *Table) Class(class string) *Table {
	t.class = class

	return t
}
