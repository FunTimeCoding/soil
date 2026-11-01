package field_row

func (r *FieldRow) Note(text string) *FieldRow {
	r.note = text

	return r
}
