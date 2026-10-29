package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listChoiceSets(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.ChoiceSets()

	if e != nil {
		return s.captureFail(e, "choice sets not listed")
	}

	return response.SuccessAny(convert.ChoiceSets(result))
}
