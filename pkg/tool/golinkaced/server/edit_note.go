package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) EditNote(
	_ context.Context,
	r server.EditNoteRequestObject,
) (server.EditNoteResponseObject, error) {
	result, f := s.client.UpdateNote(int(r.Identifier), r.Body.Text)

	if f != nil {
		return server.EditNote500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.EditNote200JSONResponse{
		Identifier: int32(result.Identifier),
		Text:       result.Text,
	}, nil
}
