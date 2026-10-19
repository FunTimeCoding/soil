package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainerSnapshots(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainerSnapshots,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	snapshots, e := s.service.ListContainerSnapshots(
		instance,
		a.Identifier,
		a.Node,
	)

	if e != nil {
		if not_found.Is(e) {
			return s.captureDetail(e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(snapshots)
}
