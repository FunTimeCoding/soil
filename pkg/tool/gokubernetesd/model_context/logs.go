package model_context

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/ambiguous_pods"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) Logs(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Logs,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	if a.Namespace == "" {
		return response.Fail("namespace is required")
	}

	cluster, e := s.activeClusterName(x)

	if e != nil {
		return response.Fail(e.Error())
	}

	q := service.LogsQuery{
		Name:       a.Name,
		Namespace:  a.Namespace,
		Container:  a.Container,
		Tail:       a.Tail,
		Since:      a.Since,
		Previous:   a.Previous,
		Timestamps: a.Timestamps,
	}

	if a.All {
		logs, f := s.service.LogsAll(x, cluster, q)

		if f != nil {
			return s.captureFail(f, "get logs")
		}

		var sections []string

		for _, l := range logs {
			sections = append(
				sections,
				fmt.Sprintf("--- %s ---\n%s", l.Pod, l.Content),
			)
		}

		return response.Success(join.NewLine(sections))
	}

	result, f := s.service.Logs(x, cluster, q)

	if f != nil {
		var ap *ambiguous_pods.AmbiguousPods

		if errors.As(f, &ap) {
			return response.Fail(f.Error())
		}

		if strings.Contains(
			f.Error(),
			"a container name must be specified",
		) {
			return response.Fail(f.Error())
		}

		return s.captureFail(f, "get logs")
	}

	return response.Success(result)
}
