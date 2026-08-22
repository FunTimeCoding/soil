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
	project, e := s.resolveProject(a.Project)

	if e != nil {
		if not_found.Is(e) {
			return response.Fail(e.Error())
		}

		return s.captureDetail(e)
	}

	result, f := s.client.SearchIssues(
		s.organization,
		a.Query,
		project,
		a.Limit,
		a.Cursor,
	)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(result)
}
