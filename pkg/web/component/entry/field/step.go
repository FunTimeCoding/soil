package field

func (f *Field) Step(value string) *Field {
	f.step = value

	return f
}
