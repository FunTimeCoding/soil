package field_row

func (r *FieldRow) Identifier(identifier string) *FieldRow {
	r.identifier = identifier

	return r
}
