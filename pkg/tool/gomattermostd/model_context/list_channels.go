package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
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

	all, e := s.client.AllChannels()

	if e != nil {
		return s.captureDetail(e)
	}

	start := a.Page * limit

	if start >= len(all) {
		start = len(all)
	}

	end := start + limit

	if end > len(all) {
		end = len(all)
	}

	page := all[start:end]
	type row struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Purpose     string `json:"purpose"`
		Header      string `json:"header"`
	}
	rows := make([]row, len(page))

	for i, c := range page {
		rows[i] = row{
			ID:          c.Id,
			Name:        c.Name,
			DisplayName: c.DisplayName,
			Type:        string(c.Type),
			Purpose:     c.Purpose,
			Header:      c.Header,
		}
	}

	return response.SuccessAny(
		map[string]any{
			"channels":    rows,
			"total_count": len(all),
			"page":        a.Page,
			"per_page":    limit,
		},
	)
}
