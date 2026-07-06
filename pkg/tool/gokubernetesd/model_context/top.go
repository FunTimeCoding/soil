package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Top(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Top,
) (*mcp.CallToolResult, error) {
	if a.ResourceType == "" {
		return response.Fail("resourceType is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	q := service.TopQuery{
		Namespace:     a.Namespace,
		AllNamespaces: a.AllNamespaces,
	}

	switch a.ResourceType {
	case "nodes":
		result, f := s.service.TopNodes(x, cluster)

		if f != nil {
			return s.captureFail(f, "top nodes")
		}

		return response.SuccessAny(result)
	case "pods":
		if a.Containers {
			result, f := s.service.TopPodContainers(x, cluster, q)

			if f != nil {
				return s.captureFail(f, "top pod containers")
			}

			return response.SuccessAny(result)
		}

		result, f := s.service.TopPods(x, cluster, q)

		if f != nil {
			return s.captureFail(f, "top pods")
		}

		return response.SuccessAny(result)
	default:
		return response.Fail(
			"unsupported resource type for top: %s (use nodes or pods)",
			a.ResourceType,
		)
	}
}
