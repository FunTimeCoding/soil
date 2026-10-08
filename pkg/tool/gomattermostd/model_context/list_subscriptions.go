package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSubscriptions(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSubscriptions,
) (*mcp.CallToolResult, error) {
	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	result, e := s.store.ByCallsign(a.Callsign)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]*mattermostResponse.Subscription, len(result))

	for i, v := range result {
		rows[i] = mattermostResponse.NewSubscription(
			v.RootIdentifier,
			v.Label(),
			v.ChannelIdentifier,
			formatTime(v.LastEvent),
		)
	}

	return response.SuccessAny(map[string]any{"subscriptions": rows})
}
