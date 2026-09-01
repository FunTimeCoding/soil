package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchEntries(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Search,
) (*mcp.CallToolResult, error) {
	return response.SuccessAny(describeCredentials(s.service.Search(a.Query)))
}
