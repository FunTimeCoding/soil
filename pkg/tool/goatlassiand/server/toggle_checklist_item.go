package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ToggleChecklistItem(
	x context.Context,
	r server.ToggleChecklistItemRequestObject,
) (server.ToggleChecklistItemResponseObject, error) {
	items, e := s.service.ToggleChecklistItem(r.Key, r.Index)

	if e != nil {
		if not_found.Is(e) {
			return server.ToggleChecklistItem404JSONResponse(
				*clientError(e),
			), nil
		}

		if ambiguous.Is(e) || not_selected.Is(e) {
			return server.ToggleChecklistItem400JSONResponse(
				*clientError(e),
			), nil
		}

		return server.ToggleChecklistItem500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ToggleChecklistItem200JSONResponse(
		convertChecklist(items),
	), nil
}
