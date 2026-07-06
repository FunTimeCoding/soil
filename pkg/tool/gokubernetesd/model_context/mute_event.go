package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) MuteEvent(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.MuteEvent,
) (*mcp.CallToolResult, error) {
	if a.Reason == "" {
		return response.Fail("reason is required")
	}

	e := s.service.MuteEvent(a.Reason, a.Message, a.Cluster)

	if e != nil {
		return s.captureFail(e, "create mute rule")
	}

	if a.Cluster != "" {
		return response.Success(
			"muted %s for cluster %s",
			a.Reason,
			a.Cluster,
		)
	}

	return response.Success("muted %s for all clusters", a.Reason)
}
