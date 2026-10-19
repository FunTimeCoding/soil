package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchIssues(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchIssues,
) (*mcp.CallToolResult, error) {
	result, e := s.service.SearchIssues(a.Query, a.Project, a.Limit, a.Cursor)

	if not_found.Is(e) {
		return response.Fail(e.Error())
	}

	if e != nil {
		return s.answer(e)
	}

	return s.succeed(result)
}
