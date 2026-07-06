package model_context_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) ReadResource(uri string) *mcp.ReadResourceResult {
	c.t.Helper()
	result, e := c.session.ReadResource(
		c.context,
		&mcp.ReadResourceParams{URI: uri},
	)
	assert.FatalOnError(c.t, e)

	return result
}
