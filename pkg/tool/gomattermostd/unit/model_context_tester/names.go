package model_context_tester

import "github.com/modelcontextprotocol/go-sdk/mcp"

func Names(v []*mcp.Tool) []string {
	var result []string

	for _, e := range v {
		result = append(result, e.Name)
	}

	return result
}
