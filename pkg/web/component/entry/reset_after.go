package entry

func (e *Entry) ResetAfter() *Entry {
	e.reset = true

	return e
}
