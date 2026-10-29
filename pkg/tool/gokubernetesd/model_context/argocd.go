package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Argocd(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Argocd,
) (*mcp.CallToolResult, error) {
	cluster, e := s.activeClusterName(x)

	if e != nil {
		return s.captureDetail(e)
	}

	if a.Name != "" {
		result, f := s.service.ArgocdApplication(
			x,
			cluster,
			a.Name,
			a.Unfiltered,
		)

		if f != nil {
			return s.captureDetail(f)
		}

		return s.formatMarkup(result.Application, result.Filtered)
	}

	result, f := s.service.ArgocdApplications(x, cluster)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(result)
}
