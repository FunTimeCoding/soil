package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Get(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Get,
) (*mcp.CallToolResult, error) {
	if a.ResourceType == "" {
		return response.Fail("resourceType is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	if a.Name != "" {
		result, f := s.service.GetResource(
			x,
			cluster,
			service.GetQuery{
				ResourceType: a.ResourceType,
				Name:         a.Name,
				Namespace:    a.Namespace,
				Unfiltered:   a.Unfiltered,
			},
		)

		if f != nil {
			return s.captureFail(f, "get resource")
		}

		return formatMarkup(result.Object, result.Filtered)
	}

	result, f := s.service.ListResources(
		x,
		cluster,
		service.ListQuery{
			ResourceType:  a.ResourceType,
			Namespace:     a.Namespace,
			AllNamespaces: a.AllNamespaces,
			LabelSelector: a.LabelSelector,
			FieldSelector: a.FieldSelector,
		},
	)

	if f != nil {
		return s.captureFail(f, "list resources")
	}

	return response.SuccessAny(result)
}
