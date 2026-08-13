package model_context_client

import "github.com/modelcontextprotocol/go-sdk/mcp"

func (c *Client) CallToolRaw(
	name string,
	arguments map[string]any,
) (*mcp.CallToolResult, error) {
	return c.session.CallTool(
		c.context,
		&mcp.CallToolParams{Name: name, Arguments: arguments},
	)
}
