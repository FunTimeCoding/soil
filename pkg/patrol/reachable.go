package patrol

import "context"

func (c *Client) Reachable(_ context.Context) error {
	s, e := c.dial()

	if e != nil {
		return e
	}

	s.Close()

	return nil
}
