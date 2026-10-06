package basic

import "github.com/funtimecoding/soil/pkg/strings/join"

func (c *Client) Next(path string) string {
	return join.Empty(c.root.String(), path)
}
