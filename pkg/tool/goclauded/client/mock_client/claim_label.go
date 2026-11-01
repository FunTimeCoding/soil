package mock_client

func (c *Client) ClaimLabel(
	sessionIdentifier string,
	key string,
	value string,
	_ string,
) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return false, c.Failure
	}

	s := c.find(sessionIdentifier)

	if s == nil || s.Labels[key] != "" {
		return false, nil
	}

	s.Labels[key] = value

	return true, nil
}
