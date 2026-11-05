package refusal

func (t *Table) Class(status int) string {
	return t.rows[status]
}
