package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateInterface(
	_ context.Context,
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

	if forbidden.Is(e) {
		return server.CreateInterface403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateInterface500JSONResponse(*s.captureDetail(e)), nil
	}

	result := server.Interface{Identifier: i.Identifier, Name: i.Name}

	if i.Type != "" {
		t := string(i.Type)
		result.Type = &t
	}

	return server.CreateInterface201JSONResponse(result), nil
}
