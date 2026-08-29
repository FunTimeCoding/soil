package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMachineSnapshots(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMachineSnapshots,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	snapshots, e := s.service.ListMachineSnapshots(
		instance,
		a.Identifier,
		a.Node,
	)

	if e != nil {
		if not_found.Is(e) {
			return response.Fail("%s", e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
