package edit_session

func (o *Session) WithTopic(value string) *Session {
	o.Topic = new(value)

	return o
}
