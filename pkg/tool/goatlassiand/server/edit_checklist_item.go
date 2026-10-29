package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) EditChecklistItem(
	x context.Context,
	r server.EditChecklistItemRequestObject,
) (server.EditChecklistItemResponseObject, error) {
	items, e := s.service.EditChecklistItem(r.Key, r.Index, r.Body.Text)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.EditChecklistItem400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.EditChecklistItem403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.EditChecklistItem404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.EditChecklistItem503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.EditChecklistItem500JSONResponse(*s.answer(e)), nil
	}

	return server.EditChecklistItem200JSONResponse(convertChecklist(items)), nil
}
