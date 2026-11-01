package entry

func (e *Entry) Inline() *Entry {
	e.inline = true

	return e
}
