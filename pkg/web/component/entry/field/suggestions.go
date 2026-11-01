package field

func (f *Field) Suggestions(values []string) *Field {
	f.suggestions = values

	return f
}
