package field

func (f *Field) Accept(types string) *Field {
	f.accept = types

	return f
}
