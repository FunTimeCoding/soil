package model_context_tester

import "github.com/mark3labs/mcp-go/mcp"

type DiscoverResponse struct {
	Result *mcp.DiscoverResult `json:"result"`
}
