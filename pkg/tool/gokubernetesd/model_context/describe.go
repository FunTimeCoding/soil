package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Describe(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Describe,
) (*mcp.CallToolResult, error) {
	if a.ResourceType == "" {
		return response.Fail("resourceType is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	result, f := s.service.DescribeResource(
		x,
		cluster,
		service.DescribeQuery{
			ResourceType: a.ResourceType,
			Name:         a.Name,
			Namespace:    a.Namespace,
			Unfiltered:   a.Unfiltered,
		},
	)

	if f != nil {
		return s.captureFail(f, "describe resource")
	}

	m := map[string]any{
		"resource": result.Resource,
		"events":   result.Events,
	}

	return formatMarkup(m, result.Filtered)
}
