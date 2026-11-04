package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SubscribeThread(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SubscribeThread,
) (*mcp.CallToolResult, error) {
	if a.Root == "" {
		return response.Fail("root is required")
	}

	if a.Callsign == "" {
		return response.Fail("callsign is required")
	}

	d := goclaudedConstant.Delivery(a.Delivery)

	if d == "" {
		d = goclaudedConstant.DeliveryImmediate
	}

	if d != goclaudedConstant.DeliveryImmediate && d != goclaudedConstant.DeliveryQueued {
		return response.Fail(
			"delivery must be immediate or queued: %s",
			a.Delivery,
		)
	}

	p, e := s.service.FindPost(a.Root)

	if e != nil {
		return s.answer(x, e)
	}

	v, existing, f := s.service.Subscribe(p, a.Callsign, a.Alias, d)

	if f != nil {
		return s.answer(x, f)
	}

	if existing {
		return response.Success("already subscribed")
	}

	result := map[string]any{
		"root":    v.RootIdentifier,
		"channel": v.ChannelIdentifier,
		"alias":   v.Alias,
	}
	replies, g := s.service.ThreadOf(v.RootIdentifier)

	if g != nil {
		result[constant.LatestReplyError] = g.Error()
	} else if len(replies) > 0 {
		result[constant.LatestReply] = replyRow(replies[len(replies)-1])
	}

	return response.SuccessAny(result)
}
