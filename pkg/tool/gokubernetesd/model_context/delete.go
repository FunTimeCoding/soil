package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Delete(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Delete,
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

	f := s.service.DeleteResource(
		x,
		cluster,
		a.ResourceType,
		a.Name,
		a.Namespace,
	)

	if f != nil {
		return s.captureFail(f, "delete resource")
	}

	namespace := a.Namespace

	if namespace == "" {
		namespace = "default"
	}

	return response.Success(
		"deleted %s/%s in %s",
		a.ResourceType,
		a.Name,
		namespace,
	)
}
