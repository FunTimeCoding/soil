package basic

import "net/url"

func (c *Client) Get(
	path string,
	query map[string]string,
	out any,
) error {
	values := url.Values{}

	for k, v := range query {
		values.Set(k, v)
	}

	return c.GetValues(path, values, out)
}
