package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchEvents(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchEvents,
) (*mcp.CallToolResult, error) {
	result, e := s.service.SearchEvents(a.Query, a.Project, a.Limit, a.Cursor)

	if not_found.Is(e) {
		return response.Fail(e.Error())
	}

	if e != nil {
		return s.captureDetail(e)
	}

	return s.succeed(result)
}
