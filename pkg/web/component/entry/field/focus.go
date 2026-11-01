package field

func (f *Field) Focus() *Field {
	f.focus = true

	return f
}
