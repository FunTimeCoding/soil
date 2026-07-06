package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListMachines(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListMachines,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	machines, e := s.service.ListMachines(c, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.Machine, len(machines))

	for i, m := range machines {
		rows[i] = proxResponse.Machine{
			Identifier: uint64(m.VMID),
			Name:       m.Name,
			Node:       m.Node,
			Status:     m.Status,
			CPU:        m.CPU,
			CPUs:       m.CPUs,
			Mem:        m.Mem,
			MaxMem:     m.MaxMem,
			Uptime:     m.Uptime,
			Tags:       m.Tags,
		}
	}

	return response.SuccessAny(rows)
}
