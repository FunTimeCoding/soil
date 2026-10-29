package client

import "context"

func (c *Client) Restart(
	x context.Context,
	name string,
) error {
	r, e := c.client.RestartProcessWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("restart", r.HTTPResponse, r.Body)
}
