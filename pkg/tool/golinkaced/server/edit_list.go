package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) EditList(
	_ context.Context,
	r server.EditListRequestObject,
) (server.EditListResponseObject, error) {
	patch := map[string]any{}

	if r.Body.Name != nil {
		patch["name"] = *r.Body.Name
	}

	if r.Body.Description != nil {
		patch["description"] = *r.Body.Description
	}

	result, f := s.client.UpdateList(int(r.Identifier), patch)

	if not_found.Is(f) {
		return server.EditList404JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.EditList500JSONResponse(*s.answer(f)), nil
	}

	return server.EditList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
