package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetOversize(
	x context.Context,
	r server.GetOversizeRequestObject,
) (server.GetOversizeResponseObject, error) {
	var collection string

	if r.Params.Collection != nil {
		collection = *r.Params.Collection
	}

	report, e := s.service.Oversize(collection)

	if e != nil {
		return server.GetOversize500JSONResponse(*s.answer(x, e)), nil
	}

	files := make([]server.OversizeFile, len(report.Files))

	for i, f := range report.Files {
		files[i] = convertOversizeFile(f)
	}

	return server.GetOversize200JSONResponse{
		Model:          report.Model,
		Window:         report.Window,
		Allowance:      report.Allowance,
		EmbedAllowance: report.EmbedAllowance,
		Files:          files,
	}, nil
}
