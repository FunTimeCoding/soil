package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/request"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) ExecInPod(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ExecInPod,
) (*mcp.CallToolResult, error) {
	cluster, e := s.activeClusterName(x)

	if e != nil {
		return s.captureDetail(e)
	}

	result, f := s.service.ExecInPod(
		x,
		cluster,
		request.Exec{
			Name:      a.Name,
			Namespace: a.Namespace,
			Command:   a.Command,
			Container: a.Container,
			Timeout:   time.Duration(a.Timeout) * time.Millisecond,
		},
	)

	if f != nil {
		return s.captureDetail(f)
	}

	if result.Failed != nil {
		if result.Stderr != "" {
			return response.Fail(
				"command failed: %s\nstderr: %s",
				result.Failed,
				result.Stderr,
			)
		}

		return s.captureFail(result.Failed, "exec in pod")
	}

	if result.Stderr != "" {
		return response.Success(
			fmt.Sprintf("%s\n--- stderr ---\n%s", result.Stdout, result.Stderr),
		)
	}

	return response.Success(result.Stdout)
}
