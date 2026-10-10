package model_context

import (
	"context"
	library "github.com/funtimecoding/soil/pkg/constant"
	markResponse "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context/response"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) summary(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	o := summary_option.New()
	o.Since = r.GetString(constant.Since, "")
	o.Until = r.GetString(constant.Until, "")
	o.GroupBy = r.GetString(constant.GroupBy, constant.Tool)
	o.Tool = r.GetString(constant.Tool, "")
	o.Actor = r.GetString(constant.Actor, "")
	rows, e := s.service.Summary(o)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	entries := make([]response.Summary, len(rows))

	for i, row := range rows {
		entries[i] = response.Summary{
			Tool:    row.Tool,
			Surface: row.Surface,
			Kind:    row.Kind,
			Outcome: row.Outcome,
			Count:   row.Count,
		}
	}

	return markResponse.SuccessAny(entries)
}
