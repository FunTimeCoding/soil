package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateSnippet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateSnippet,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.Content == "" {
		return response.Fail("content is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	c.WriteFile(snippetPath(a.Name), []byte(a.Content))

	return response.SuccessAny(
		map[string]string{
			"volume": fmt.Sprintf("local:snippets/%s", a.Name),
			"status": "created",
		},
	)
}
