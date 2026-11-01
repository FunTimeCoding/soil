package field

func (f *Field) Required() *Field {
	f.required = true

	return f
}
