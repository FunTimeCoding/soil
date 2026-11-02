package gitlab

func (c *Client) userIdentifier() (int64, error) {
	c.mutex.Lock()
	u := c.user
	c.mutex.Unlock()

	if u != nil {
		return u.ID, nil
	}

	u, e := c.CurrentUser()

	if e != nil {
		return 0, e
	}

	c.mutex.Lock()
	c.user = u
	c.mutex.Unlock()

	return u.ID, nil
}
