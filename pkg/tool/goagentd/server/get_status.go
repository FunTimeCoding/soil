package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	status := s.runner.Status()

	return server.GetStatus200JSONResponse{
		State:  string(status.State),
		Result: &status.Result,
	}, nil
}
