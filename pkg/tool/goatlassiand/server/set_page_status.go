package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SetPageStatus(
	_ context.Context,
	r server.SetPageStatusRequestObject,
) (server.SetPageStatusResponseObject, error) {
	if r.Body.Status != "current" && r.Body.Status != "draft" {
		return server.SetPageStatus400JSONResponse(
			*clientError(
				not_selected.Format(
					"status must be 'current' (publish) or 'draft' (unpublish)",
				),
			),
		), nil
	}

	result, e := s.service.SetPageStatus(r.Identifier, r.Body.Status)

	if e != nil {
		return server.SetPageStatus500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetPageStatus200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
