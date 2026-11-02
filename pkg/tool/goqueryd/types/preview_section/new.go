package preview_section

func New(
	level int,
	title string,
	firstLine int,
	lastLine int,
	tokens int,
) *Section {
	return &Section{
		Level:     level,
		Title:     title,
		FirstLine: firstLine,
		LastLine:  lastLine,
		Tokens:    tokens,
	}
}
