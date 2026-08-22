package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/provision/model_context"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) runs(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	limit := r.GetInt(constant.Limit, 20)
	result, e := s.store.Recent(limit)

	if e != nil {
		return s.captureFail(e, generativeConstant.RecentRunsFail)
	}

	summaries := make([]model_context.RunSummary, len(result))

	for i, v := range result {
		summaries[i] = model_context.RunSummary{
			ID:                  v.ID,
			CreatedAt:           v.CreatedAt.Format(timeConstant.DateSecond),
			Scope:               v.Scope,
			TriggerSource:       v.TriggerSource,
			DurationMillisecond: v.DurationMillisecond,
			Status:              v.Status,
			GitHead:             v.GitHead,
		}
	}

	return response.SuccessAny(summaries)
}
