package entry

func (e *Entry) Query() *Entry {
	e.query = true

	return e
}
