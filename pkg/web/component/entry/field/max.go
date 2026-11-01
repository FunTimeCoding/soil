package field

func (f *Field) Max(value string) *Field {
	f.maximum = value

	return f
}
