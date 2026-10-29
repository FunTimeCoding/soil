package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) captureFail(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	return s.answerWith(e, message)
}
