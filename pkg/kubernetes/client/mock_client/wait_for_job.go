package mock_client

import "time"

func (c *Client) WaitForJob(_ string, _ string, _ time.Duration) error {
	return c.Failure
}
