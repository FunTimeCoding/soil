package edit_session

func (o *Session) WithAlias(value string) *Session {
	o.Alias = new(value)

	return o
}
