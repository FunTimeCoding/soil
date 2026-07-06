package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/provision/model_context"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) runs(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	limit := r.GetInt(constant.Limit, 20)
	result, e := s.store.Recent(limit)

	if e != nil {
		return s.captureFail(e, constant.RecentRunsFailed)
	}

	summaries := make([]model_context.RunSummary, len(result))

	for i, v := range result {
		summaries[i] = model_context.RunSummary{
			ID:                  v.ID,
			CreatedAt:           v.CreatedAt.Format(time.DateSecond),
			Scope:               v.Scope,
			TriggerSource:       v.TriggerSource,
			DurationMillisecond: v.DurationMillisecond,
			Status:              v.Status,
			GitHead:             v.GitHead,
		}
	}

	return response.SuccessAny(summaries)
}
