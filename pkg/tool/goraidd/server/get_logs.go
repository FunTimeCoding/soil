package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
)

func (s *Server) GetLogs(
	_ context.Context,
	r server.GetLogsRequestObject,
) (server.GetLogsResponseObject, error) {
	all := s.store.Fights()
	var fights []raid.Fight

	for _, f := range all {
		if r.Params.Start != nil && f.Timestamp.Before(*r.Params.Start) {
			continue
		}

		if r.Params.End != nil && f.Timestamp.After(*r.Params.End) {
			continue
		}

		fights = append(fights, f)
	}

	total := len(fights)
	offset := 0
	limit := 50

	if r.Params.Offset != nil {
		offset = *r.Params.Offset
	}

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	if offset > total {
		offset = total
	}

	end := offset + limit

	if end > total {
		end = total
	}

	page := fights[offset:end]
	result := make([]server.LogResponse, 0, len(page))

	for _, f := range page {
		result = append(
			result,
			server.LogResponse{
				FileName:    f.Filename,
				Time:        f.Timestamp.Format(time.DateSecond),
				Duration:    fmt.Sprintf("%dms", f.DurationMS),
				MapId:       f.MapID,
				PlayerCount: f.AlliedCount,
				Players:     nil,
			},
		)
	}

	return server.GetLogs200JSONResponse{
		Total: total,
		Logs:  result,
	}, nil
}
