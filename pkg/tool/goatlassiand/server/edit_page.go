package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) EditPage(
	_ context.Context,
	r server.EditPageRequestObject,
) (server.EditPageResponseObject, error) {
	title := ""
	message := ""
	draft := false

	if r.Body.Title != nil {
		title = *r.Body.Title
	}

	if r.Body.Message != nil {
		message = *r.Body.Message
	}

	if r.Body.Draft != nil {
		draft = *r.Body.Draft
	}

	result, e := s.service.EditPage(
		r.Identifier,
		r.Body.OldText,
		r.Body.NewText,
		title,
		message,
		draft,
	)

	if validation.Is(e) || ambiguous.Is(e) {
		return server.EditPage400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.EditPage403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.EditPage404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.EditPage503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.EditPage500JSONResponse(*s.answer(e)), nil
	}

	return server.EditPage200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
