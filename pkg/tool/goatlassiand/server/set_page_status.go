package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SetPageStatus(
	_ context.Context,
	r server.SetPageStatusRequestObject,
) (server.SetPageStatusResponseObject, error) {
	if r.Body.Status != "current" && r.Body.Status != "draft" {
		return server.SetPageStatus400JSONResponse{
			Error: not_selected.Format(
				"status must be 'current' (publish) or 'draft' (unpublish)",
			).Error(),
		}, nil
	}

	result, e := s.service.SetPageStatus(r.Identifier, r.Body.Status)

	if forbidden.Is(e) {
		return server.SetPageStatus403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.SetPageStatus404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.SetPageStatus503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SetPageStatus500JSONResponse(*s.answer(e)), nil
	}

	return server.SetPageStatus200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
