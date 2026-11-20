package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) PutSystem(
	x context.Context,
	r server.PutSystemRequestObject,
) (server.PutSystemResponseObject, error) {
	e := s.system.Put(convert.Snapshot(r.Name, r.Body))

	if validation.Is(e) {
		return server.PutSystem400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PutSystem500JSONResponse(*s.answer(x, e)), nil
	}

	return server.PutSystem204Response{}, nil
}
