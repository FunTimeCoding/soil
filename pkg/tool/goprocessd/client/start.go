package client

import "context"

func (c *Client) Start(
	x context.Context,
	name string,
) error {
	r, e := c.client.StartProcessWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("start", r.HTTPResponse, r.Body)
}
