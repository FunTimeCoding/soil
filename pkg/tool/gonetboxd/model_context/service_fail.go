package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) serviceFail(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return s.answerWith(e, message)
}
