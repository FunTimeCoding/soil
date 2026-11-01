package field

func (f *Field) Lines(count string) *Field {
	f.lines = count

	return f
}
