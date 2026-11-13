package highlight

func NewMarked(
	text string,
	open string,
	close string,
) *Highlight {
	return &Highlight{text: text, open: open, close: close}
}
