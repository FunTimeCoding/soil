package model_context

import "github.com/modelcontextprotocol/go-sdk/mcp"

func toolNames(v []*mcp.Tool) []string {
	var result []string

	for _, e := range v {
		result = append(result, e.Name)
	}

	return result
}
