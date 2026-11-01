package field_row

func (r *FieldRow) Value(text string) *FieldRow {
	r.value = text

	return r
}
