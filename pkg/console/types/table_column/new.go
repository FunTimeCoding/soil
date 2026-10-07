package table_column

func New(
	header string,
	width int,
) *Column {
	return &Column{Header: header, Width: width}
}
