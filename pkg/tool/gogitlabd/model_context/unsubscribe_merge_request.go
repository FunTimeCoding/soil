package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UnsubscribeMergeRequest(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UnsubscribeMergeRequest,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	removed, f := s.service.UnsubscribeMergeRequest(
		instance,
		project,
		a.MergeRequest,
		a.Callsign,
	)

	if f != nil {
		return s.answerMessage(x, f, "")
	}

	if !removed {
		return response.Success("no such subscription")
	}

	return response.Success("unsubscribed")
}
