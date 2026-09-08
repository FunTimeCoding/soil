package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
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

	type row struct {
		Root      string `json:"root"`
		Label     string `json:"label"`
		Channel   string `json:"channel"`
		LastEvent string `json:"last_event"`
	}
	rows := make([]row, len(result))

	for i, v := range result {
		rows[i] = row{
			Root:      v.RootIdentifier,
			Label:     v.Label(),
			Channel:   v.ChannelIdentifier,
			LastEvent: formatTime(v.LastEvent),
		}
	}

	return response.SuccessAny(map[string]any{"subscriptions": rows})
}
