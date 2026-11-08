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

	if s == nil || len(s.Labels[key]) > 0 {
		return false, nil
	}

	s.Labels[key] = []string{value}

	return true, nil
}
