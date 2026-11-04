package mock_client

func (c *Client) FailFor(
	callsign string,
	e error,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failures == nil {
		c.failures = map[string]error{}
	}

	if e == nil {
		delete(c.failures, callsign)

		return
	}

	c.failures[callsign] = e
}
