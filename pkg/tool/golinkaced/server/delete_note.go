package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteNote(
	_ context.Context,
	r server.DeleteNoteRequestObject,
) (server.DeleteNoteResponseObject, error) {
	if e := s.client.DeleteNote(int(r.Identifier)); e != nil {
		return server.DeleteNote500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteNote200Response{}, nil
}
