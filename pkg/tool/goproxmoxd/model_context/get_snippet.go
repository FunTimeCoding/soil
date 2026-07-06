package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetSnippet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetSnippet,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	b := c.ReadFile(snippetPath(a.Name))

	return response.SuccessAny(
		map[string]string{
			"name":    a.Name,
			"content": string(b),
		},
	)
}
