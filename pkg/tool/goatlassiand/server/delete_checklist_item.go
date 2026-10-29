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

func (s *Server) DeleteChecklistItem(
	x context.Context,
	r server.DeleteChecklistItemRequestObject,
) (server.DeleteChecklistItemResponseObject, error) {
	items, e := s.service.DeleteChecklistItem(r.Key, r.Index)

	if validation.Is(e) || ambiguous.Is(e) || not_selected.Is(e) {
		return server.DeleteChecklistItem400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.DeleteChecklistItem403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.DeleteChecklistItem404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeleteChecklistItem503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteChecklistItem500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteChecklistItem200JSONResponse(
		convertChecklist(items),
	), nil
}
