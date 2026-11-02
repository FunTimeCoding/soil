package result

func New(
	columns []string,
	rows [][]any,
) *Result {
	return &Result{Columns: columns, Rows: rows}
}
