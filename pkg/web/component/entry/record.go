package entry

func (e *Entry) Record() *Entry {
	e.record = true

	return e
}
