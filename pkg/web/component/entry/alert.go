package entry

func (e *Entry) Alert(message string) *Entry {
	e.alert = message

	return e
}
