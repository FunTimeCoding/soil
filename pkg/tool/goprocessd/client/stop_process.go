package client

import "context"

func (c *Client) StopProcess(
	x context.Context,
	name string,
) error {
	r, e := c.client.StopProcessWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("stop", r.HTTPResponse, r.Body)
}
