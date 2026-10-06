package model_context

import "github.com/mark3labs/mcp-go/mcp"

func textOr(
	q mcp.CallToolRequest,
	name string,
	stored string,
) string {
	if value := q.GetString(name, ""); value != "" {
		return value
	}

	return stored
}
