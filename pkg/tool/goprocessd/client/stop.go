package client

import "context"

func (c *Client) Stop(
	x context.Context,
	name string,
) error {
	r, e := c.client.StopProcessWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("stop", r.HTTPResponse, r.Body)
}
