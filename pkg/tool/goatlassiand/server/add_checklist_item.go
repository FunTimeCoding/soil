package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) AddChecklistItem(
	x context.Context,
	r server.AddChecklistItemRequestObject,
) (server.AddChecklistItemResponseObject, error) {
	items, e := s.service.AddChecklistItem(r.Key, r.Body.Text)

	if e != nil {
		if not_found.Is(e) {
			return server.AddChecklistItem404JSONResponse(*clientError(e)), nil
		}

		if ambiguous.Is(e) || not_selected.Is(e) {
			return server.AddChecklistItem400JSONResponse(*clientError(e)), nil
		}

		return server.AddChecklistItem500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.AddChecklistItem200JSONResponse(convertChecklist(items)), nil
}
