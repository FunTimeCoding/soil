package mock_client

func (c *Client) Equip(key string) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.gear.Equipped["weapon"] = key

	return nil
}
