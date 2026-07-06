package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Certificates(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Certificates,
) (*mcp.CallToolResult, error) {
	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	if a.Name != "" {
		if a.Namespace == "" {
			return response.Fail("namespace is required for detail view")
		}

		result, f := s.service.Certificate(
			x,
			cluster,
			a.Name,
			a.Namespace,
			a.Unfiltered,
		)

		if f != nil {
			return s.captureFail(f, "get certificate")
		}

		m := map[string]any{
			"certificate": result.Certificate,
		}

		if result.LatestRequest != nil {
			m["latestRequest"] = result.LatestRequest
		}

		return formatMarkup(m, result.Filtered)
	}

	result, f := s.service.Certificates(x, cluster)

	if f != nil {
		return s.captureFail(f, "list certificates")
	}

	return response.SuccessAny(result)
}
