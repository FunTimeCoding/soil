package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteContainerSnapshot(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteContainerSnapshot,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	taskIdentifier, e := s.service.DeleteContainerSnapshot(
		instance,
		a.Identifier,
		a.Node,
		a.Name,
	)

	if e != nil {
		if not_found.Is(e) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"task_id": taskIdentifier})
}
