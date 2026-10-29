package client

import "context"

func (c *Client) ClearLog(
	x context.Context,
	name string,
) error {
	r, e := c.client.ClearLogWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("clear log", r.HTTPResponse, r.Body)
}
