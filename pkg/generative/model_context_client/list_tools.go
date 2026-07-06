package model_context_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) ListTools() []*mcp.Tool {
	c.t.Helper()
	result, e := c.session.ListTools(
		c.context,
		&mcp.ListToolsParams{},
	)
	assert.FatalOnError(c.t, e)

	return result.Tools
}
