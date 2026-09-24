package table

func New(headers ...string) *Table {
	result := &Table{}

	for _, h := range headers {
		result.columns = append(
			result.columns,
			&column{header: h, width: len(h)},
		)
	}

	return result
}
