package mock_client

import "slices"

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

	if s == nil || !slices.Contains(s.Labels[key], value) {
		return false, nil
	}

	s.Labels[key] = slices.DeleteFunc(
		s.Labels[key],
		func(v string) bool { return v == value },
	)

	if len(s.Labels[key]) == 0 {
		delete(s.Labels, key)
	}

	return true, nil
}
