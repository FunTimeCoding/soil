package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"sort"
	"time"
)

func (s *Server) GetCost(
	_ context.Context,
	r server.GetCostRequestObject,
) (server.GetCostResponseObject, error) {
	days := 1

	if r.Params.Days != nil && *r.Params.Days > 0 {
		days = *r.Params.Days
	}

	to := time.Now().UTC()
	from := to.Add(-time.Duration(days) * 24 * time.Hour)
	usage := s.service.UsageBetween(from, to)
	result := server.CostResponse{}

	if len(usage) == 0 {
		return server.GetCost200JSONResponse(result), nil
	}

	models := make([]string, 0, len(usage))

	for model := range usage {
		models = append(models, model)
	}

	sort.Strings(models)
	var rows []server.ModelUsage

	for _, model := range models {
		t := usage[model]
		cost := pricing.Cost(model, t)
		result.Cost += cost
		rows = append(
			rows,
			server.ModelUsage{
				Model:         model,
				Calls:         t.Count,
				Input:         t.Input,
				Output:        t.Output,
				CacheCreation: t.CacheCreation,
				CacheRead:     t.CacheRead,
				Cache5Minute:  t.Cache5Minute,
				Cache1Hour:    t.Cache1Hour,
				Cost:          cost,
			},
		)
	}

	result.Models = &rows

	return server.GetCost200JSONResponse(result), nil
}
