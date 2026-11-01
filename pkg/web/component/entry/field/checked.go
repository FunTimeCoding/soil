package field

func (f *Field) Checked(checked bool) *Field {
	f.checked = checked

	return f
}
