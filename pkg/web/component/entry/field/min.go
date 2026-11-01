package field

func (f *Field) Min(value string) *Field {
	f.minimum = value

	return f
}
