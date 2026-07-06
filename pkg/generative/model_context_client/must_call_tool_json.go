package model_context_client

import "github.com/funtimecoding/soil/pkg/assert"

func (c *Client) MustCallToolJSON(
	name string,
	arguments map[string]any,
) map[string]any {
	c.t.Helper()
	result, e := c.CallToolJSON(name, arguments)
	assert.FatalOnError(c.t, e)

	return result
}
