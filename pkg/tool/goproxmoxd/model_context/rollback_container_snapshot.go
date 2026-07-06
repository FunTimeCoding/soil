package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RollbackContainerSnapshot(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RollbackContainerSnapshot,
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

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	taskID, e := s.service.RollbackContainerSnapshot(
		c,
		a.Identifier,
		a.Node,
		a.Name,
	)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(map[string]string{"task_id": taskID})
}
