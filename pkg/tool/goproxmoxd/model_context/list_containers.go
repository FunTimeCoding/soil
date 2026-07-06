package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	proxResponse "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListContainers(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListContainers,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	containers, e := s.service.ListContainers(c, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]proxResponse.Container, len(containers))

	for i, ct := range containers {
		rows[i] = proxResponse.Container{
			Identifier: uint64(ct.VMID),
			Name:       ct.Name,
			Node:       ct.Node,
			Status:     ct.Status,
			CPUs:       ct.CPUs,
			MaxMem:     ct.MaxMem,
			MaxDisk:    ct.MaxDisk,
			MaxSwap:    ct.MaxSwap,
			Uptime:     ct.Uptime,
			Tags:       ct.Tags,
		}
	}

	return response.SuccessAny(rows)
}
