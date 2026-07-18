package telegram

func (c *Client) saveChannel(
	identifier int64,
	name string,
) {
	if c.store == nil {
		return
	}

	c.store.MustSaveChannel(identifier, name)
}
