package field

func (f *Field) Placeholder(hint string) *Field {
	f.placeholder = hint

	return f
}
