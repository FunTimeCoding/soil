package edit_session

func (o *Session) WithFiles(value string) *Session {
	o.Files = new(value)

	return o
}
