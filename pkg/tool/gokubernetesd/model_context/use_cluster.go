package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UseCluster(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UseCluster,
) (*mcp.CallToolResult, error) {
	if a.Cluster == "" {
		return response.Fail("cluster is required")
	}

	if _, e := s.service.ClusterByName(a.Cluster); e != nil {
		return s.captureDetail(e)
	}

	if e := s.selection.Select(x, a.Cluster); e != nil {
		return s.captureDetail(e)
	}

	return response.Success("active cluster set to %s", a.Cluster)
}
