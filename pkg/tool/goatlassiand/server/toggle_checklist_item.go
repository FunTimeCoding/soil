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

func (s *Server) ToggleChecklistItem(
	x context.Context,
	r server.ToggleChecklistItemRequestObject,
) (server.ToggleChecklistItemResponseObject, error) {
	items, e := s.service.ToggleChecklistItem(r.Key, r.Index)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.ToggleChecklistItem400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.ToggleChecklistItem403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.ToggleChecklistItem404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ToggleChecklistItem503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ToggleChecklistItem500JSONResponse(*s.answer(e)), nil
	}

	return server.ToggleChecklistItem200JSONResponse(
		convertChecklist(items),
	), nil
}
