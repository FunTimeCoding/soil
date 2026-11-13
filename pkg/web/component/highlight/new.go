package highlight

func New(
	text string,
	terms ...string,
) *Highlight {
	return &Highlight{text: text, terms: terms}
}
