package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) GetAffectedSystems(
	x context.Context,
	r server.GetAffectedSystemsRequestObject,
) (server.GetAffectedSystemsResponseObject, error) {
	result, e := s.system.Affected(r.Identifier)

	if validation.Is(e) {
		return server.GetAffectedSystems400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetAffectedSystems500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetAffectedSystems200JSONResponse(
		*convert.AffectedSystems(result),
	), nil
}
