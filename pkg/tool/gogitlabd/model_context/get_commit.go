package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetCommit(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetCommit,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Sha == "" {
		return response.Fail("sha is required")
	}

	v, _, e := s.client.Commits.GetCommit(a.Project, a.Sha, nil)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
