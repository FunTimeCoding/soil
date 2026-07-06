package model_context_client

import "github.com/funtimecoding/soil/pkg/assert"

func (c *Client) MustCallToolError(
	name string,
	arguments map[string]any,
) string {
	c.t.Helper()
	result, e := c.CallToolError(name, arguments)
	assert.FatalOnError(c.t, e)

	return result
}
