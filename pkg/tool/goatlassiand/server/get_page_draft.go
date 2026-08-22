package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPageDraft(
	_ context.Context,
	r server.GetPageDraftRequestObject,
) (server.GetPageDraftResponseObject, error) {
	result, e := s.confluence.DraftOverlay(r.Identifier)

	if e != nil {
		return server.GetPageDraft500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetPageDraft200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
