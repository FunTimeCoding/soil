package model_context_client

import "github.com/funtimecoding/soil/pkg/assert"

func (c *Client) MustCallToolNotation(
	name string,
	arguments map[string]any,
) map[string]any {
	c.t.Helper()
	result, e := c.CallToolNotation(name, arguments)
	assert.FatalOnError(c.t, e)

	return result
}
