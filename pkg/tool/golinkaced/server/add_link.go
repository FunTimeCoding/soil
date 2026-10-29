package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AddLink(
	_ context.Context,
	r server.AddLinkRequestObject,
) (server.AddLinkResponseObject, error) {
	value := ""

	if r.Body.List != nil {
		value = *r.Body.List
	}

	listIdentifier, e := s.service.ResolveList(value)

	if validation.Is(e) {
		return server.AddLink400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.AddLink403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.AddLink503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AddLink500JSONResponse(*s.answer(e)), nil
	}

	name := ""

	if r.Body.Name != nil {
		name = *r.Body.Name
	}

	if name == "" {
		name = r.Body.Link
	}

	result, f := s.service.CreateLink(r.Body.Link, name, listIdentifier, nil)

	if validation.Is(f) {
		return server.AddLink400JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.AddLink403JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.AddLink503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.AddLink500JSONResponse(*s.answer(f)), nil
	}

	return server.AddLink200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Title,
		Link:       result.Link,
	}, nil
}
