package line_diff

func NewAnchor(
	line int,
	rewritten bool,
) *Anchor {
	return &Anchor{Line: line, Rewritten: rewritten}
}
