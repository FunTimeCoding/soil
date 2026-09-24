package changed

func New(
	base string,
	head string,
) *Range {
	return &Range{Base: base, Head: head}
}
