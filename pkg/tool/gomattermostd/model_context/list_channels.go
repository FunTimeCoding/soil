package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	mattermostResponse "github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListChannels(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListChannels,
) (*mcp.CallToolResult, error) {
	limit := a.Limit

	if limit <= 0 {
		limit = 100
	}

	page, e := s.client.AllChannels(limit, a.Page*limit)

	if e != nil {
		return s.captureDetail(e)
	}

	rows := make([]*mattermostResponse.Channel, len(page))

	for i, c := range page {
		rows[i] = mattermostResponse.NewChannel(
			c.Id,
			c.Name,
			c.DisplayName,
			channelTypeName(c.Type),
			c.Purpose,
			c.Header,
		)
	}

	return response.SuccessAny(
		map[string]any{
			"channels": rows,
			"page":     a.Page,
			"per_page": limit,
			"has_more": len(rows) == limit,
		},
	)
}
