package table

func New(columns ...string) *Table {
	return &Table{columns: columns, numeric: map[int]bool{}}
}
