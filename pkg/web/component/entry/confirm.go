package entry

func (e *Entry) Confirm(message string) *Entry {
	e.confirm = message

	return e
}
