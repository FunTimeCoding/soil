package statement

func New(
	sql string,
	a []any,
) *Statement {
	return &Statement{SQL: sql, Arguments: a}
}
