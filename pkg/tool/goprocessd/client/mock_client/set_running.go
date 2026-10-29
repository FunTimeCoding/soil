package mock_client

func (c *Client) setRunning(
	name string,
	running bool,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	i, e := c.index(name)

	if e != nil {
		return e
	}

	c.ProcessList[i].Running = running

	return nil
}
