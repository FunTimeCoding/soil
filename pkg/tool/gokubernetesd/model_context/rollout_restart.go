package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RolloutRestart(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.RolloutRestart,
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

	f := s.service.RolloutRestart(
		x,
		cluster,
		a.ResourceType,
		a.Name,
		a.Namespace,
	)

	if f != nil {
		return s.captureFail(f, "rollout restart")
	}

	namespace := a.Namespace

	if namespace == "" {
		namespace = "default"
	}

	return response.Success(
		"restarted %s/%s in %s",
		a.ResourceType,
		a.Name,
		namespace,
	)
}
