package refusing

import "net/url"

func (c *Client) URL(
	endpoint string,
	arguments map[string]string,
) *url.URL {
	return c.client.URL(endpoint, arguments)
}
