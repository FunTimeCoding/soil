package technitium

import "net/url"

func (c *Client) post(
	path string,
	form url.Values,
) error {
	_, e := c.basic.Post(path, form)

	return e
}
