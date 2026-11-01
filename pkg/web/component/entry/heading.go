package entry

func (e *Entry) Heading(text string) *Entry {
	e.heading = text

	return e
}
