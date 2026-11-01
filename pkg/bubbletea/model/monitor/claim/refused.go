package claim

func (c *Connection) Refused(e error) {
	c.updates <- Message{Error: e}
}
