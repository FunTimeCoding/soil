package entry

func (e *Entry) Sideways() *Entry {
	e.sideways = true

	return e
}
