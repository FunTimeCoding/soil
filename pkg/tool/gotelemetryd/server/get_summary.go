package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
)

func (s *Server) GetSummary(
	_ context.Context,
	r server.GetSummaryRequestObject,
) (server.GetSummaryResponseObject, error) {
	since := ""
	until := ""
	groupBy := "tool"

	if r.Params.Since != nil {
		since = *r.Params.Since
	}

	if r.Params.Until != nil {
		until = *r.Params.Until
	}

	if r.Params.GroupBy != nil {
		groupBy = string(*r.Params.GroupBy)
	}

	rows, e := s.store.Summary(since, until, groupBy)

	if e != nil {
		return server.GetSummary500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	entries := make([]server.SummaryEntry, 0, len(rows))

	for _, r := range rows {
		entry := server.SummaryEntry{
			Tool:  r.Tool,
			Count: int(r.Count),
		}

		if r.Surface != "" {
			entry.Surface = &r.Surface
		}

		if r.Kind != "" {
			entry.Kind = &r.Kind
		}

		entries = append(entries, entry)
	}

	return server.GetSummary200JSONResponse(entries), nil
}
