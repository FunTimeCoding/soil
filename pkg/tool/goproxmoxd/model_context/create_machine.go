package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a create_machine.Machine,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
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

	identifier, e := s.service.CreateMachine(c, &a)

	if e != nil {
		if errors.Is(e, constant.ErrorCDROMCloudInitConflict) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier": identifier,
			"status":     "created",
		},
	)
}
