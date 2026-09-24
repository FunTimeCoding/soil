package table

func (t *Table) headers() []string {
	result := make([]string, len(t.columns))

	for i, c := range t.columns {
		result[i] = c.header
	}

	return result
}
