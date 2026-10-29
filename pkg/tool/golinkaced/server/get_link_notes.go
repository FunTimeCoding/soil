package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) GetLinkNotes(
	_ context.Context,
	r server.GetLinkNotesRequestObject,
) (server.GetLinkNotesResponseObject, error) {
	page := 1

	if r.Params.Page != nil {
		page = int(*r.Params.Page)
	}

	v, e := s.service.NotesByLinkPage(int(r.Identifier), page)

	if not_found.Is(e) {
		return server.GetLinkNotes404JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.GetLinkNotes403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetLinkNotes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetLinkNotes500JSONResponse(*s.answer(e)), nil
	}

	result := make([]server.Note, len(v.Items))

	for i, n := range v.Items {
		result[i] = server.Note{
			Identifier: int32(n.Identifier),
			Text:       n.Text,
		}
	}

	total := v.Total

	return server.GetLinkNotes200JSONResponse(
		server.NotePage{Total: &total, Notes: &result},
	), nil
}
