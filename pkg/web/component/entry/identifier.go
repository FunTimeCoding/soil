package entry

func (e *Entry) Identifier(identifier string) *Entry {
	e.identifier = identifier

	return e
}
