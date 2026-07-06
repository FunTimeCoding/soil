package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

func (c *Client) PostPath(
	path string,
	body string,
) (int, string, error) {
	return c.Post(locator.New(c.host).Path(path).String(), body)
}
