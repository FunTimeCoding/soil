package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) MitigateFlag(
	x context.Context,
	r server.MitigateFlagRequestObject,
) (server.MitigateFlagResponseObject, error) {
	mitigation := ""

	if r.Body.Mitigation != nil {
		mitigation = *r.Body.Mitigation
	}

	result, e := s.flag.Mitigate(r.Identifier, mitigation)

	if validation.Is(e) {
		return server.MitigateFlag400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.MitigateFlag404JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.MitigateFlag409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.MitigateFlag500JSONResponse(*s.answer(x, e)), nil
	}

	return server.MitigateFlag200JSONResponse(*convert.Flag(result)), nil
}
