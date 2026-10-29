package mock_client

import "github.com/funtimecoding/soil/pkg/strings/join"

func (c *Client) record(v ...string) {
	c.calls = append(c.calls, join.Space(v...))
}
