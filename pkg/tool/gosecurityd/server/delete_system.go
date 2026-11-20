package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) DeleteSystem(
	x context.Context,
	r server.DeleteSystemRequestObject,
) (server.DeleteSystemResponseObject, error) {
	e := s.system.Delete(r.Name)

	if not_found.Is(e) {
		return server.DeleteSystem404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteSystem500JSONResponse(*s.answer(x, e)), nil
	}

	return server.DeleteSystem204Response{}, nil
}
