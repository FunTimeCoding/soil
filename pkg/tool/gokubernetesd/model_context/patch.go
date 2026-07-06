package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Patch(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Patch,
) (*mcp.CallToolResult, error) {
	if a.ResourceType == "" {
		return response.Fail("resourceType is required")
	}

	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.Patch == "" {
		return response.Fail("patch is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	f := s.service.PatchResource(
		x,
		cluster,
		service.PatchQuery{
			ResourceType: a.ResourceType,
			Name:         a.Name,
			Namespace:    a.Namespace,
			Patch:        a.Patch,
			Type:         resolvePatchType(a.Type),
			DryRun:       a.DryRun,
		},
	)

	if f != nil {
		return s.captureFail(f, "patch resource")
	}

	namespace := a.Namespace

	if namespace == "" {
		namespace = "default"
	}

	prefix := "patched"

	if a.DryRun {
		prefix = "dry run: would patch"
	}

	return response.Success(
		"%s %s/%s in %s",
		prefix,
		a.ResourceType,
		a.Name,
		namespace,
	)
}
