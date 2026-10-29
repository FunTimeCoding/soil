package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) ReloadEnvironment(
	_ context.Context,
	_ server.ReloadEnvironmentRequestObject,
) (server.ReloadEnvironmentResponseObject, error) {
	if e := s.service.ReloadEnvironment(); e != nil {
		return server.ReloadEnvironment500JSONResponse(*s.answer(e)), nil
	}

	return server.ReloadEnvironment204Response{}, nil
}
