package client

import "context"

func (c *Client) RemoveDeviceLabel(
	x context.Context,
	device string,
	key string,
) error {
	r, e := c.client.RemoveDeviceLabelWithResponse(x, device, key)

	if e != nil {
		return e
	}

	return noContent("remove device label", r.HTTPResponse, r.Body)
}
