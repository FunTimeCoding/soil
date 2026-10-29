package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) OpenFile(
	_ context.Context,
	r server.OpenFileRequestObject,
) (server.OpenFileResponseObject, error) {
	v, e := s.service.OpenFile(r.Body.FilePath)

	if validation.Is(e) {
		return server.OpenFile400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.OpenFile500JSONResponse(*s.answer(e)), nil
	}

	return server.OpenFile200JSONResponse(*convertView(v)), nil
}
