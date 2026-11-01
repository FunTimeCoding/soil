package entry

func (e *Entry) EndGroup() *Entry {
	return e.Group("")
}
