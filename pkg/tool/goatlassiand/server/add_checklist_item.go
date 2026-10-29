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

func (s *Server) AddChecklistItem(
	x context.Context,
	r server.AddChecklistItemRequestObject,
) (server.AddChecklistItemResponseObject, error) {
	items, e := s.service.AddChecklistItem(r.Key, r.Body.Text)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.AddChecklistItem400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.AddChecklistItem403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.AddChecklistItem404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.AddChecklistItem503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AddChecklistItem500JSONResponse(*s.answer(e)), nil
	}

	return server.AddChecklistItem200JSONResponse(convertChecklist(items)), nil
}
