package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	return s.answerWith(e, "")
}
