package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteProjectVariable(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteProjectVariable,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Key == "" {
		return response.Fail("key is required")
	}

	project, e := s.resolveProject(a.Project)

	if e != nil {
		return s.captureDetail(e)
	}

	e = s.client.DeleteProjectVariable(project, a.Key)

	if e != nil {
		return s.captureDetail(e)
	}

	return mcp.NewToolResultText("deleted"), nil
}
