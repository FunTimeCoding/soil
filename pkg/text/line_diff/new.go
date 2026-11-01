package line_diff

func New(
	mark string,
	text string,
) *Line {
	return &Line{Mark: mark, Text: text}
}
