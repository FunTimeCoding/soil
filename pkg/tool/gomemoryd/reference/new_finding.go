package reference

func NewFinding(
	span string,
	text string,
) *Finding {
	return &Finding{Span: span, Text: text}
}
