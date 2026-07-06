package model_context_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) ListResources() []*mcp.Resource {
	c.t.Helper()
	result, e := c.session.ListResources(
		c.context,
		&mcp.ListResourcesParams{},
	)
	assert.FatalOnError(c.t, e)

	return result.Resources
}
