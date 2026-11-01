package field_row

func (r *FieldRow) Tooltip(text string) *FieldRow {
	r.tooltip = text

	return r
}
