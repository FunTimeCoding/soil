package client

import "context"

func (c *Client) DeleteCable(
	x context.Context,
	identifier int32,
) error {
	r, e := c.client.DeleteCableWithResponse(x, identifier)

	if e != nil {
		return e
	}

	return noContent("delete cable", r.HTTPResponse, r.Body)
}
