package mock_client

import "strconv"

func (c *Client) DeleteCable(identifier int32) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.record("DeleteCable", strconv.Itoa(int(identifier)))

	return nil
}
