package model_context

import (
	"context"
	mark "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListClusters(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListClusters,
) (*mcp.CallToolResult, error) {
	active, e := s.selection.Active(x)

	if e != nil {
		return s.captureDetail(e)
	}

	result := []response.ClusterEntry{}

	for _, c := range s.service.Clusters() {
		result = append(
			result,
			response.ClusterEntry{Name: c.Name(), Active: c.Name() == active},
		)
	}

	return mark.SuccessAny(result)
}
