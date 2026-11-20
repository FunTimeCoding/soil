package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) GetDebianStatus(
	x context.Context,
	r server.GetDebianStatusRequestObject,
) (server.GetDebianStatusResponseObject, error) {
	release := ""

	if r.Params.Release != nil {
		release = *r.Params.Release
	}

	result, e := s.debian.Status(r.Identifier, release)

	if validation.Is(e) {
		return server.GetDebianStatus400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetDebianStatus404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetDebianStatus503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetDebianStatus500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetDebianStatus200JSONResponse(
		*convert.DebianStatus(result),
	), nil
}
