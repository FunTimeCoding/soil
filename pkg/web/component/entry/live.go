package entry

func (e *Entry) Live() *Entry {
	e.live = true

	return e
}
