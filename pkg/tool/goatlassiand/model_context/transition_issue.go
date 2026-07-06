package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) transitionIssue(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	transition, g := r.RequireString(constant.TransitionIdentifier)

	if g != nil {
		return response.Fail("transition_identifier is required: %v", g)
	}

	if h := s.jira.Transition(key, transition); h != nil {
		return s.captureFail(h, "transition not applied")
	}

	return response.Success("transition applied")
}
