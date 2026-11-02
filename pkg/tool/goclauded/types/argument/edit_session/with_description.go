package edit_session

func (o *Session) WithDescription(value string) *Session {
	o.Description = new(value)

	return o
}
