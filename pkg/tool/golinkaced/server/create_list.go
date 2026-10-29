package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) CreateList(
	_ context.Context,
	r server.CreateListRequestObject,
) (server.CreateListResponseObject, error) {
	description := ""

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	result, f := s.service.CreateList(r.Body.Name, description)

	if validation.Is(f) {
		return server.CreateList400JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.CreateList403JSONResponse{Error: f.Error()}, nil
	}

	if conflict.Is(f) {
		return server.CreateList409JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.CreateList503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.CreateList500JSONResponse(*s.answer(f)), nil
	}

	return server.CreateList200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
