package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a update_machine.Machine,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	e = s.service.UpdateMachine(instance, &a)

	if e != nil {
		if not_found.Is(e) || validation.Is(e) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	return response.Success("updated")
}
