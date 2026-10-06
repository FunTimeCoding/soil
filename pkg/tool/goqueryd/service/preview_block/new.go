package preview_block

func New(
	kind string,
	firstLine int,
	lastLine int,
	tokens int,
) *Block {
	return &Block{
		Kind:      kind,
		FirstLine: firstLine,
		LastLine:  lastLine,
		Tokens:    tokens,
	}
}
