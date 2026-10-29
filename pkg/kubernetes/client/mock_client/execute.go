package mock_client

import "io"

func (c *Client) Execute(_ io.Writer, _ io.Writer, _ string, _ string, _ string, _ ...string) error {
	return c.Failure
}
