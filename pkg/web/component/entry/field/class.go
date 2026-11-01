package field

func (f *Field) Class(class string) *Field {
	f.class = class

	return f
}
