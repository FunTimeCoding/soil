package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	count, e := s.service.Count()

	if e != nil {
		return server.GetStatus500JSONResponse(*s.answer(e)), nil
	}

	return server.GetStatus200JSONResponse{TotalEntries: count}, nil
}
