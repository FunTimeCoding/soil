package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SubscribeThread(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SubscribeThread,
) (*mcp.CallToolResult, error) {
	if a.Root == "" {
		return response.Fail("root is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	p, e := s.service.FindPost(a.Root)

	if e != nil {
		return s.answer(e)
	}

	v, existing, f := s.service.Subscribe(p, a.Callsign, a.Alias)

	if f != nil {
		return s.answer(f)
	}

	if existing {
		return response.Success("already subscribed")
	}

	return response.SuccessAny(
		map[string]any{
			"root":    v.RootIdentifier,
			"channel": v.ChannelIdentifier,
			"alias":   v.Alias,
		},
	)
}
