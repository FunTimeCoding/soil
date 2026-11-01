package mock_client

func (c *Client) ReleaseLabel(
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

	if c.ReleaseFailure != nil {
		return false, c.ReleaseFailure
	}

	s := c.find(sessionIdentifier)

	if s == nil || s.Labels[key] != value {
		return false, nil
	}

	delete(s.Labels, key)

	return true, nil
}
