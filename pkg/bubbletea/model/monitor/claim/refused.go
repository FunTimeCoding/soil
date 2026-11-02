package claim

func (c *Connection) Refused(e error) {
	c.deliver(Message{Error: e})
}
