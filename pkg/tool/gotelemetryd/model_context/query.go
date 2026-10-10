package model_context

import (
	"context"
	"encoding/json"
	library "github.com/funtimecoding/soil/pkg/constant"
	markResponse "github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context/response"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) query(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	o := query_option.New()
	o.Tool = r.GetString(constant.Tool, "")
	o.Surface = r.GetString(constant.Surface, "")
	o.Actor = r.GetString(constant.Actor, "")
	o.Kind = r.GetString(constant.Kind, "")
	o.Since = r.GetString(constant.Since, "")
	o.Until = r.GetString(constant.Until, "")
	o.Limit = r.GetInt(constant.Limit, 50)
	events, e := s.service.Query(o)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	entries := make([]response.Event, len(events))

	for i, v := range events {
		entries[i] = response.Event{
			Tool:      v.Tool,
			Surface:   v.Surface,
			Actor:     v.Actor,
			Outcome:   v.Outcome,
			Kind:      v.Kind,
			Duration:  v.DurationMillisecond,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
		}
		var detail map[string]string

		if v.Detail != nil && json.Unmarshal([]byte(*v.Detail), &detail) == nil {
			entries[i].Detail = detail
		}
	}

	return markResponse.SuccessAny(entries)
}
