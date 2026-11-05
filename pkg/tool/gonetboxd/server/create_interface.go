package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateInterface(
	x context.Context,
	r server.CreateInterfaceRequestObject,
) (server.CreateInterfaceResponseObject, error) {
	address := ""

	if r.Body.PhysicalAddress != nil {
		address = *r.Body.PhysicalAddress
	}

	i, e := s.service.CreateInterface(r.Name, r.Body.Name, r.Body.Type, address)

	if not_found.Is(e) {
		return server.CreateInterface404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreateInterface400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateInterface503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateInterface500JSONResponse(*s.captureDetail(x, e)), nil
	}

	result := server.Interface{Identifier: i.Identifier, Name: i.Name}

	if i.Type != "" {
		t := string(i.Type)
		result.Type = &t
	}

	return server.CreateInterface201JSONResponse(result), nil
}
