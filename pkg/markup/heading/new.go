package heading

func New(
	text string,
	slug string,
) *Heading {
	return &Heading{Text: text, Slug: slug}
}
