package timestamp_column

func New(
	table string,
	name string,
) *Column {
	return &Column{Table: table, Name: name}
}
