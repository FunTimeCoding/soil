package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
)

func (s *Server) GetSummary(
	_ context.Context,
	r server.GetSummaryRequestObject,
) (server.GetSummaryResponseObject, error) {
	o := summary_option.New()

	if r.Params.Since != nil {
		o.Since = *r.Params.Since
	}

	if r.Params.Until != nil {
		o.Until = *r.Params.Until
	}

	if r.Params.GroupBy != nil {
		o.GroupBy = string(*r.Params.GroupBy)
	}

	rows, e := s.store.Summary(o)

	if e != nil {
		return server.GetSummary500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	entries := make([]server.SummaryEntry, 0, len(rows))

	for _, r := range rows {
		entry := server.SummaryEntry{Tool: r.Tool, Count: int(r.Count)}

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
