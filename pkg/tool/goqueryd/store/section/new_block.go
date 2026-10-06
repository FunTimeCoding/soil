package section

func NewBlock(
	kind string,
	firstLine int,
	lastLine int,
	text string,
) *Block {
	return &Block{
		Kind:      kind,
		FirstLine: firstLine,
		LastLine:  lastLine,
		Text:      text,
	}
}
