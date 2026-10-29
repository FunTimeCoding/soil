package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/server"
	"time"
)

func (s *Server) PostGenerate(
	_ context.Context,
	r server.PostGenerateRequestObject,
) (server.PostGenerateResponseObject, error) {
	date := time.Now()

	if r.Body.Date != nil {
		date = *r.Body.Date
	}

	v, e := s.service.Generate(r.Body.Files, date)

	if not_found.Is(e) {
		return server.PostGenerate404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostGenerate500JSONResponse(*s.answer(e)), nil
	}

	return server.PostGenerate200JSONResponse{
		Path:     v.Path,
		TidCount: v.TidCount,
	}, nil
}
