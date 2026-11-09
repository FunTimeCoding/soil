package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
)

func (s *Server) GetLogs(
	x context.Context,
	r server.GetLogsRequestObject,
) (server.GetLogsResponseObject, error) {
	w := window.Whole()

	if r.Params.Start != nil {
		w.WithStart(*r.Params.Start)
	}

	if r.Params.End != nil {
		w.WithEnd(*r.Params.End)
	}

	var fights []fight.Fight

	for _, f := range s.service.Fights() {
		if w.Contains(f.Timestamp) {
			fights = append(fights, f)
		}
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
				Time:        f.Timestamp.Format(timeConstant.DateSecond),
				Duration:    fmt.Sprintf("%dms", f.DurationMS),
				MapId:       f.MapIdentifier,
				PlayerCount: f.AlliedCount,
				Players:     nil,
			},
		)
	}

	return server.GetLogs200JSONResponse{Total: total, Logs: result}, nil
}
