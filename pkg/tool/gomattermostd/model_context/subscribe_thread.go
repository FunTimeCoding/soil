package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
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

	p, e := s.client.FindPost(a.Root)

	if e != nil {
		return s.captureFail(e, "post not found")
	}

	root := p.Id

	if p.RootId != "" {
		root = p.RootId
	}

	existing, f := s.store.ByRoot(root)

	if f != nil {
		return s.captureDetail(f)
	}

	for _, v := range existing {
		if v.Callsign == a.Callsign {
			return response.Success("already subscribed")
		}
	}

	if g := s.store.Create(
		subscription.New(a.Callsign, root, p.ChannelId, a.Alias),
	); g != nil {
		return s.captureDetail(g)
	}

	s.indexer.Index(root)

	return response.SuccessAny(
		map[string]any{"root": root, "channel": p.ChannelId, "alias": a.Alias},
	)
}
