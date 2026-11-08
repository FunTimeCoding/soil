package mock_client

import "slices"

func (c *Client) AddLabel(
	sessionIdentifier string,
	key string,
	value string,
	_ string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	s := c.find(sessionIdentifier)

	if s == nil || slices.Contains(s.Labels[key], value) {
		return nil
	}

	s.Labels[key] = append(s.Labels[key], value)

	return nil
}
