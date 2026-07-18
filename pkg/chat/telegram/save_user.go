package telegram

func (c *Client) saveUser(
	identifier int64,
	name string,
) {
	if c.store == nil {
		return
	}

	c.store.MustSaveUser(identifier, name)
}
