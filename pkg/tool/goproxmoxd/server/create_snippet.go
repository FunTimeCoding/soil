package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateSnippet(
	x context.Context,
	r server.CreateSnippetRequestObject,
) (server.CreateSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.CreateSnippet404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.CreateSnippet400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateSnippet500JSONResponse(*s.answer(x, e)), nil
	}

	volume, e := s.service.CreateSnippet(instance, r.Body.Name, r.Body.Content)

	switch {
	case validation.Is(e):
		return server.CreateSnippet400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.CreateSnippet404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.CreateSnippet403JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.CreateSnippet500JSONResponse(*s.answer(x, e)), nil
	}

	return server.CreateSnippet200JSONResponse{
		Volume: volume,
		Status: "created",
	}, nil
}
