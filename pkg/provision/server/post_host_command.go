package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
)

func (s *Server) PostHostCommand(
	x context.Context,
	r server.PostHostCommandRequestObject,
) (server.PostHostCommandResponseObject, error) {
	var host string

	if r.Body.Host != nil {
		host = *r.Body.Host
	}

	result, e := s.runner.HostCommand(r.Body.Command, host, r.Body.Source)

	if validation.Is(e) {
		return server.PostHostCommand400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.PostHostCommand404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.PostHostCommand503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostHostCommand500JSONResponse(*s.answer(x, e)), nil
	}

	return server.PostHostCommand200JSONResponse(hostCommandRun(result)), nil
}
