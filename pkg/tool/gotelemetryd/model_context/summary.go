package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	markResponse "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) summary(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	since := r.GetString(constant.Since, "")
	until := r.GetString(constant.Until, "")
	groupBy := r.GetString(constant.GroupBy, constant.Tool)
	rows, e := s.service.Summary(since, until, groupBy)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	entries := make([]response.Summary, len(rows))

	for i, row := range rows {
		entries[i] = response.Summary{
			Tool:    row.Tool,
			Surface: row.Surface,
			Kind:    row.Kind,
			Count:   row.Count,
		}
	}

	return markResponse.SuccessAny(entries)
}
