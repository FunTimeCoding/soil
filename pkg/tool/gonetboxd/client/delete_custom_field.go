package client

import "context"

func (c *Client) DeleteCustomField(
	x context.Context,
	name string,
) error {
	r, e := c.client.DeleteCustomFieldWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("delete custom field", r.HTTPResponse, r.Body)
}
