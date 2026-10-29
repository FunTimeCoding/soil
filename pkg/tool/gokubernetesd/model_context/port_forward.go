package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) PortForward(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.PortForward,
) (*mcp.CallToolResult, error) {
	cluster, e := s.activeClusterName(x)

	if e != nil {
		return s.captureDetail(e)
	}

	f, g := s.service.PortForward(
		x,
		cluster,
		request.PortForward{
			ResourceType: a.ResourceType,
			Name:         a.Name,
			Namespace:    a.Namespace,
			LocalPort:    a.LocalPort,
			TargetPort:   a.TargetPort,
		},
	)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(
		map[string]any{
			"id":         f.Identifier,
			"localPort":  f.LocalPort,
			"targetPort": f.TargetPort,
			"resource":   f.Resource,
			"namespace":  f.Namespace,
		},
	)
}
