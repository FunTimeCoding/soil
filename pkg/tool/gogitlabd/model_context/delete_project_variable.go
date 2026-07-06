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

	_, e := s.client.ProjectVariables.RemoveVariable(
		a.Project,
		a.Key,
		nil,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return mcp.NewToolResultText("deleted"), nil
}
