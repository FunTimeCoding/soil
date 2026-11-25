package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	gitlabResponse "github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSubscriptions(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSubscriptions,
) (*mcp.CallToolResult, error) {
	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	result, e := s.service.MergeRequestSubscriptions(a.Callsign)

	if e != nil {
		return s.answerMessage(x, e, "")
	}

	rows := make([]*gitlabResponse.Subscription, len(result))

	for i, v := range result {
		rows[i] = gitlabResponse.NewSubscription(v)
	}

	return response.SuccessAny(map[string]any{"subscriptions": rows})
}
