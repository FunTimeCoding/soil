package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/luthermonson/go-proxmox"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CloneMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CloneMachine,
) (*mcp.CallToolResult, error) {
	if a.Identifier == 0 {
		return response.Fail("identifier is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	options := &proxmox.VirtualMachineCloneOptions{Name: a.Name}

	if a.NewIdentifier > 0 {
		options.NewID = a.NewIdentifier
	}

	if a.Full {
		options.Full = true
	}

	if a.Storage != "" {
		options.Storage = a.Storage
	}

	if a.Snapshot != "" {
		options.SnapName = a.Snapshot
	}

	newIdentifier, e := s.service.CloneMachine(
		instance,
		a.Identifier,
		a.Node,
		options,
	)

	if e != nil {
		if not_found.Is(e) {
			return s.captureDetail(e)
		}

		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"identifier": newIdentifier, "status": "cloned"},
	)
}
