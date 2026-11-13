package outage_tester

import "github.com/modelcontextprotocol/go-sdk/mcp"

func text(r *mcp.CallToolResult) string {
	for _, c := range r.Content {
		if t, okay := c.(*mcp.TextContent); okay {
			return t.Text
		}
	}

	return ""
}
