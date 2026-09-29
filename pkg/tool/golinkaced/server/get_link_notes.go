package server

import (
	"context"
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

	v, e := s.client.NotesByLinkPage(int(r.Identifier), page)

	if e != nil {
		return server.GetLinkNotes500JSONResponse(*s.captureDetail(e)), nil
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
