package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetPageDraft(
	_ context.Context,
	r server.GetPageDraftRequestObject,
) (server.GetPageDraftResponseObject, error) {
	result, e := s.service.PageDraft(r.Identifier)

	if forbidden.Is(e) {
		return server.GetPageDraft403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetPageDraft404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetPageDraft503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetPageDraft500JSONResponse(*s.answer(e)), nil
	}

	return server.GetPageDraft200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
