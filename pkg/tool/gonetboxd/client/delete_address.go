package client

import "context"

func (c *Client) DeleteAddress(
	x context.Context,
	identifier int32,
) error {
	r, e := c.client.DeleteAddressWithResponse(x, identifier)

	if e != nil {
		return e
	}

	return noContent("delete address", r.HTTPResponse, r.Body)
}
