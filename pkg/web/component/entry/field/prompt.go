package field

func (f *Field) Prompt(text string) *Field {
	f.prompt = text

	return f
}
