package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) EditTag(
	_ context.Context,
	r server.EditTagRequestObject,
) (server.EditTagResponseObject, error) {
	patch := map[string]any{}

	if r.Body.Name != nil {
		patch["name"] = *r.Body.Name
	}

	result, f := s.service.UpdateTag(int(r.Identifier), patch)

	if not_found.Is(f) {
		return server.EditTag404JSONResponse{Error: f.Error()}, nil
	}

	if validation.Is(f) {
		return server.EditTag400JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.EditTag403JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.EditTag503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.EditTag500JSONResponse(*s.answer(f)), nil
	}

	return server.EditTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
