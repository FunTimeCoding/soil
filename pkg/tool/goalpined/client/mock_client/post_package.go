package mock_client

import (
	"context"
	"io"
)

func (c *Client) PostPackage(
	_ context.Context,
	_ string,
	_ string,
	_ string,
	file string,
	_ io.Reader,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.Uploaded = append(c.Uploaded, file)

	return nil
}
