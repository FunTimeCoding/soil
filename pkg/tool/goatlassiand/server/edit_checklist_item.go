package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) EditChecklistItem(
	x context.Context,
	r server.EditChecklistItemRequestObject,
) (server.EditChecklistItemResponseObject, error) {
	items, e := s.service.EditChecklistItem(r.Key, r.Index, r.Body.Text)

	if e != nil {
		if not_found.Is(e) {
			return server.EditChecklistItem404JSONResponse(*clientError(e)), nil
		}

		if ambiguous.Is(e) || not_selected.Is(e) {
			return server.EditChecklistItem400JSONResponse(*clientError(e)), nil
		}

		return server.EditChecklistItem500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.EditChecklistItem200JSONResponse(convertChecklist(items)), nil
}
