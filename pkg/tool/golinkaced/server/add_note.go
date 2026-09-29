package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) AddNote(
	_ context.Context,
	r server.AddNoteRequestObject,
) (server.AddNoteResponseObject, error) {
	result, f := s.client.CreateNote(int(r.Body.LinkIdentifier), r.Body.Text)

	if f != nil {
		return server.AddNote500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.AddNote200JSONResponse{
		Identifier: int32(result.Identifier),
		Text:       result.Text,
	}, nil
}
