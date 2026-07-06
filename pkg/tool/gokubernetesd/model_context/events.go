package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Events(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Events,
) (*mcp.CallToolResult, error) {
	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	result, f := s.service.Events(
		x,
		cluster,
		service.EventsQuery{
			Namespace:  a.Namespace,
			Kind:       a.Kind,
			Name:       a.Name,
			Type:       a.Type,
			Limit:      a.Limit,
			Unfiltered: a.Unfiltered,
		},
	)

	if f != nil {
		return s.captureFail(f, "list events")
	}

	return response.SuccessAny(result)
}
