package board_limit

func New(
	board string,
	column string,
	maximum int,
) *Limit {
	return &Limit{Board: board, Column: column, Maximum: maximum}
}
