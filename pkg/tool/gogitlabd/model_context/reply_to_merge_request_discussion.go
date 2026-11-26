package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReplyToMergeRequestDiscussion(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ReplyToMergeRequestDiscussion,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.MergeRequest == 0 {
		return response.Fail("merge_request is required")
	}

	if a.Discussion == "" {
		return response.Fail("discussion is required")
	}

	if a.Body == "" {
		return response.Fail("body is required")
	}

	instance, project, e := s.resolveProject(x, a.Project)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	v, f := s.service.ReplyToMergeRequestDiscussion(
		x,
		instance,
		project,
		a.MergeRequest,
		a.Discussion,
		a.Body,
		a.Callsign,
		a.Resolve,
	)

	if f != nil && v == nil {
		return s.answerMessage(x, f, "")
	}

	if f != nil {
		return s.answerMessage(
			x,
			f,
			"the reply was posted, but resolving the discussion failed",
		)
	}

	return response.SuccessAny(v)
}
