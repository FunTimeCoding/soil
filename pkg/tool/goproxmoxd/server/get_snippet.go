package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetSnippet(
	x context.Context,
	r server.GetSnippetRequestObject,
) (server.GetSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.GetSnippet404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.GetSnippet400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetSnippet500JSONResponse(*s.answer(x, e)), nil
	}

	b, e := s.service.GetSnippet(instance, r.Name)

	switch {
	case validation.Is(e):
		return server.GetSnippet400JSONResponse(*clientError(e)), nil
	case not_found.Is(e):
		return server.GetSnippet404JSONResponse(*clientError(e)), nil
	case forbidden.Is(e):
		return server.GetSnippet403JSONResponse(*clientError(e)), nil
	case e != nil:
		return server.GetSnippet500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetSnippet200JSONResponse{
		Name:    r.Name,
		Content: string(b),
	}, nil
}
