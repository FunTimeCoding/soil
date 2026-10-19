package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeleteChecklistItem(
	x context.Context,
	r server.DeleteChecklistItemRequestObject,
) (server.DeleteChecklistItemResponseObject, error) {
	items, e := s.service.DeleteChecklistItem(r.Key, r.Index)

	if e != nil {
		if not_found.Is(e) {
			return server.DeleteChecklistItem404JSONResponse(
				*clientError(e),
			), nil
		}

		if ambiguous.Is(e) || not_selected.Is(e) {
			return server.DeleteChecklistItem400JSONResponse(
				*clientError(e),
			), nil
		}

		return server.DeleteChecklistItem500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.DeleteChecklistItem200JSONResponse(
		convertChecklist(items),
	), nil
}
