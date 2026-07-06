package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Apply(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Apply,
) (*mcp.CallToolResult, error) {
	if a.Manifest == "" {
		return response.Fail("manifest is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	result, f := s.service.ApplyResource(
		x,
		cluster,
		service.ApplyQuery{
			Manifest:  a.Manifest,
			Namespace: a.Namespace,
			Override:  a.Override,
			DryRun:    a.DryRun,
		},
	)

	if f != nil {
		return s.captureFail(f, "apply resource")
	}

	prefix := "applied"

	if a.DryRun {
		prefix = "dry run: would apply"
	}

	return response.Success(
		fmt.Sprintf(
			"%s %s/%s in %s",
			prefix,
			result.Kind,
			result.Name,
			result.Namespace,
		),
	)
}
