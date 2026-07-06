package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListNodes(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListNodes,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	nodes, e := s.service.ListNodes(c)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.Node, len(nodes))

	for i, n := range nodes {
		rows[i] = proxResponse.Node{
			Name:    n.Node,
			Status:  n.Status,
			CPU:     n.CPU,
			MaxCPU:  n.MaxCPU,
			Mem:     n.Mem,
			MaxMem:  n.MaxMem,
			Disk:    n.Disk,
			MaxDisk: n.MaxDisk,
			Uptime:  n.Uptime,
		}
	}

	return response.SuccessAny(rows)
}
