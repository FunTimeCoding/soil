package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

func (c *Client) GetPath(path string) (int, string, error) {
	return c.Get(locator.New(c.host).Path(path).String())
}
